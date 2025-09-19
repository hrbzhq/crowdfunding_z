package main

import "testing"

func TestShouldEnableAutoupdate(t *testing.T) {
    cases := []struct{
        args []string
        env string
        want bool
    }{
        {[]string{}, "false", false},
        {[]string{"--autoupdate"}, "false", true},
        {[]string{}, "true", true},
        {[]string{"--other"}, "TRUE", true},
    }
    for i, c := range cases {
        got := shouldEnableAutoupdate(c.args, c.env)
        if got != c.want {
            t.Fatalf("case %d: got %v want %v (args=%v env=%q)", i, got, c.want, c.args, c.env)
        }
    }
}
