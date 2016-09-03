package main

import (
  "os"
  "fmt"
  "bufio"
  "syscall"
  "strings"
  "golang.org/x/crypto/ssh/terminal"
  "github.com/codegangsta/cli"
  "github.com/google/go-github/github"
)

func prompt(Message string) string {
  fmt.Print(Message)
  reader := bufio.NewReader(os.Stdin)
  text, _ := reader.ReadString('\n')
  return strings.TrimSpace(text)
}

func password(Message string) string {
  fmt.Print(Message)
  bytes, _ := terminal.ReadPassword(int(syscall.Stdin))
  text := string(bytes)
  fmt.Println()
  return strings.TrimSpace(text)
}

func main() {
  app           := cli.NewApp()
  app.Name       = "HubToken"
  app.Usage      = "Manage GitHub personal access tokens"
  app.Version    = "0.1.0"
  app.Commands   = []cli.Command{
    {
      Name: "create",
      Aliases: []string{"c"},
      Usage: "Create personal access token called `NAME`",
      Action: func(c *cli.Context) error {
        note      := c.Args().First()
        login     := prompt("Github login: ")
        password  := password("password: ")
        otp       := prompt("2FA/OTP: ")
        scopes    := []github.Scope{"repo"}
        tp        := github.BasicAuthTransport{Username: login, Password: password, OTP: otp}
        client    := github.NewClient(tp.Client())
        input     := &github.AuthorizationRequest{Note: &note, Scopes: scopes}

        authorization, _, err := client.Authorizations.Create(input)

        if err != nil {
          fmt.Errorf("Authorizations.Create returned error: %v", err)
          return nil
        }

        token := *authorization.Token
        fmt.Printf("%v\n", strings.TrimPrefix(token, "0x"))

        return nil
      },
    },
    {
      Name: "delete",
      Aliases: []string{"D"},
      Usage: "Delete a token called `NAME`",
      Action: func(c *cli.Context) error {
        // Request Authentication Parameters (Login, Password, OTP)
        // Call Github API
        fmt.Println("deleting ...", c.Args().First())
        return nil
      },
    },
    {
      Name: "list",
      Aliases: []string{"l"},
      Usage: "list personal access tokens",
      Action: func(c *cli.Context) error {
        // Request Authentication Parameters (Login, Password, OTP)
        // Call Github API
        fmt.Println("listing ...")
        return nil
      },
    },
  }

  app.Run(os.Args)
}
