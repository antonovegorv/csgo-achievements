package matchestracker

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/antonovegorv/csgo-achievements/internal/pkg/models"
	"github.com/antonovegorv/csgo-achievements/internal/pkg/store"
	"github.com/antonovegorv/csgo-achievements/internal/pkg/util"
	"github.com/rs/zerolog"
)

type client struct {
	httpClient    *http.Client
	store         store.Store
	logger        zerolog.Logger
	apiKey        string
	apiEndpoint   string
	tickerTimeout int
	loggerLevel   string
}

type response struct {
	Result struct {
		NextCode string `json:"nextcode"`
	} `json:"result"`
}

func newClient(store store.Store, apiKey, apiEndpoint string, tickerTimeout int, loggerLevel string) *client {
	c := &client{
		httpClient:    &http.Client{Timeout: time.Second * 15},
		store:         store,
		apiKey:        apiKey,
		apiEndpoint:   apiEndpoint,
		tickerTimeout: tickerTimeout,
		loggerLevel:   loggerLevel,
	}

	return c
}

func (c *client) trackMatches() error {
	if err := c.configureLogger(); err != nil {
		return err
	}

	t := time.NewTicker(time.Duration(c.tickerTimeout) * time.Second)

	for {
		c.logger.Info().Msg("grabbing all users with their last matches...")

		lums, err := c.store.LastUserMatch().GetAll()
		if err != nil {
			return err
		}

		for _, lum := range lums {
			c.logger.Info().Msgf("trying to update info about user(steam_id=%d)", lum.User.SteamID)

			if err := c.updateUser(lum); err != nil {
				return err
			}

			c.logger.Info().Msg("user has been updated")
		}

		c.logger.Info().Msgf("grabbing completed, next try in %d seconds", c.tickerTimeout)

		<-t.C
	}
}

func (c *client) configureLogger() error {
	level, err := zerolog.ParseLevel(c.loggerLevel)
	if err != nil {
		return err
	}

	c.logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		Level(level).
		With().
		Timestamp().
		Caller().
		Logger()

	return nil
}

func (c *client) updateUser(lum *models.LastUserMatch) error {
	req, err := http.NewRequest(http.MethodGet, c.apiEndpoint, nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("key", c.apiKey)
	q.Add("steamid", strconv.Itoa(lum.User.SteamID))
	q.Add("steamidkey", lum.User.GameAuthenticationCode)
	q.Add("knowncode", lum.Match.ShareCode)

	req.URL.RawQuery = q.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	c.logger.Info().Msgf("response code = %d", resp.StatusCode)

	if resp.StatusCode == http.StatusOK {
		response := &response{}
		if err = json.NewDecoder(resp.Body).Decode(response); err != nil {
			return err
		}

		c.logger.Info().Msgf("new match(share_code=%s) was founded", response.Result.NextCode)

		m, err := c.store.Match().FindBySharingCode(response.Result.NextCode)
		if err != nil {
			if err != sql.ErrNoRows {
				return err
			}

			m = util.DecodeShareCode(response.Result.NextCode)

			m, err = c.store.Match().Create(m)
			if err != nil {
				return err
			}

			c.logger.Info().Msg("new match was created")
		}

		mh := &models.MatchHistory{User: lum.User, Match: m}

		_, err = c.store.MatchHistory().Create(mh)
		if err != nil {
			return err
		}

		if err = c.store.LastUserMatch().UpdateByUserID(lum.User.ID, m.ID); err != nil {
			return err
		}
	}

	return nil
}
