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
  "github.com/fatih/color"
)

func InfoMessage() *color.Color {
  return color.New(color.FgGreen, color.Bold)
}

func HeadingMessage() *color.Color {
  return color.New(color.Bold)
}

func ErrorMessage() *color.Color {
  return color.New(color.FgRed, color.Bold)
}

func isPersonalAccessToken(a *github.Authorization) bool {
  return *a.App.URL == "https://developer.github.com/v3/oauth_authorizations/"
}

func ForEachAuthorizations(authorizations []*github.Authorization, f func(*github.Authorization)) {
  for _, authorization := range authorizations {
    if isPersonalAccessToken(authorization) {
      f(authorization)
    }
  }
}
func Prompt(Message string) string {
  fmt.Print(Message)
  reader := bufio.NewReader(os.Stdin)
  text, _ := reader.ReadString('\n')
  return strings.TrimSpace(text)
}

func PasswordPrompt(Message string) string {
  fmt.Print(Message)
  bytes, _ := terminal.ReadPassword(int(syscall.Stdin))
  text := string(bytes)
  fmt.Println()
  return strings.TrimSpace(text)
}

func Login() github.BasicAuthTransport {
  return github.BasicAuthTransport{
    Username: Prompt("Github login: "),
    Password: PasswordPrompt("Password: "),
    OTP: Prompt("2FA/OTP: "),
  }
}

func CreateToken(note string) {
  login := Login()
  client := github.NewClient(login.Client())

  // TODO: Scopes should be set by the user
  scopes := []github.Scope{"repo"}

  auth_req := &github.AuthorizationRequest{
    Note: &note,
    Scopes: scopes,
  }

  authorization, _, err := client.Authorizations.Create(auth_req)

  if err != nil {
    fmt.Printf("Error creating personal access token: %v\nAuthorizations.Create returned error: %v\n", note, err)
  } else {
    token := *authorization.Token
    fmt.Printf("%v\n", strings.TrimPrefix(token, "0x"))
  }
}

func GetAuthorizationsList(client *github.Client) []*github.Authorization {
  authorizations, _, err := client.Authorizations.List(nil)
  if err != nil {
    ErrorMessage().Printf("Error getting personal access tokens\nAuthorizations.List returned error: %v\n", err)
    os.Exit(1)
  }
  return authorizations
}


func ListTokens() {
  login := Login()
  client := github.NewClient(login.Client())
  authorizations := GetAuthorizationsList(client)
  if len(authorizations) > 0 {
    HeadingMessage().Printf("\nPersonal Access Tokens for %s:\n", login.Username)
    ForEachAuthorizations(authorizations, func(auth *github.Authorization) {
      InfoMessage().Printf("%s\n", *auth.Note)
    })
  } else {
    ErrorMessage().Printf("There are no personal access tokens for your GitHub account\n")
  }
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
        CreateToken(c.Args().First())
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
        ListTokens()
        return nil
      },
    },
  }

  app.Run(os.Args)
}
