package env

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func (c *client) RequireBool(key string) (bool, error) {
	if b, err := strconv.ParseBool(c.Getenv(key)); err == nil {
		return b, nil
	}
	return false, keyError(key)
}

// Bytes returns a slice of bytes from the ENV, or fallback variable
func (c *client) RequireBytes(key string) ([]byte, error) {
	if v := c.Getenv(key); v != "" {
		return []byte(v), nil
	}

	return nil, keyError(key)
}

func (c *client) RequireFloat64(key string) (float64, error) {
	if f, err := strconv.ParseFloat(c.Getenv(key), 64); err == nil {
		return f, nil
	}
	return 0.0, keyError(key)
}

func (c *client) RequireDuration(key string) (time.Duration, error) {
	if d, err := time.ParseDuration(c.Getenv(key)); err == nil {
		return d, nil
	}

	return 0, keyError(key)
}

func (c *client) RequireInt(key string) (int, error) {
	if i, err := strconv.Atoi(c.Getenv(key)); err == nil {
		return i, nil
	}
	return 0, keyError(key)
}

func (c *client) RequireString(key string) (string, error) {
	if v := c.Getenv(key); v != "" {
		return v, nil
	}
	return "", keyError(key)
}

// RequireStrings returns a slice of strings from the ENV
func (c *client) RequireStrings(key string, seps ...string) ([]string, error) {
	if v := c.Getenv(key); v != "" {
		sep := ","

		if len(seps) > 0 {
			sep = seps[0]
		}

		return strings.Split(v, sep), nil
	}

	return []string{}, keyError(key)
}

// URL returns a URL from the ENV, or fallback URL if missing/invalid
func (c *client) RequireURL(key string) (*url.URL, error) {
	if v := c.Getenv(key); v != "" {
		u, err := url.Parse(v)
		if err != nil {
			return u, nil
		}
		return u, nil
	}
	return nil, keyError(key)
}

func (c *client) RequireAddr(key string) (net.IP, error) {
	if v := c.Getenv(key); v != "" {
		addr := net.ParseIP(v)
		if addr != nil {
			return addr, nil
		} else {
			return nil, errors.New(fmt.Sprintf("%s is not a valid address", key))
		}
	}
	return nil, keyError(key)
}

// RequireBool does not allow defaults
func RequireBool(key string) (bool, error) {
	return DefaultClient.RequireBool(key)
}

// RequireBytes does not allow defaults
func RequireBytes(key string) ([]byte, error) {
	return DefaultClient.RequireBytes(key)
}

// RequireFloat64 does not allow defaults
func RequireFloat64(key string) (float64, error) {
	return DefaultClient.RequireFloat64(key)
}

// RequireDuration does not allow defaults
func RequireDuration(key string) (time.Duration, error) {
	return DefaultClient.RequireDuration(key)
}

// RequireInt does not allow defaults
func RequireInt(key string) (int, error) {
	return DefaultClient.RequireInt(key)
}

// RequireString does not allow defaults.
func RequireString(key string) (string, error) {
	return DefaultClient.RequireString(key)
}

// RequireStrings does not allow defaults
func RequireStrings(key string) ([]string, error) {
	return DefaultClient.RequireStrings(key)
}

// RequireURL does not allow defaults
func RequireURL(key string) (*url.URL, error) {
	return DefaultClient.RequireURL(key)
}

// RequireAddr does not allow defaults
func RequireAddr(key string) (net.IP, error) {
	return DefaultClient.RequireAddr(key)
}

// KeyError returns an error string
func keyError(key string) error {
	return errors.New(key + " is a requried environment variable.")
}

// Required is a fatal checking function. Exiting from a library is frowned upon
func Required(e error) {
	if e != nil {
		panic(fmt.Sprintf("Fatal error: %s", e.Error()))
	}
}
