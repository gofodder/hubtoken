package main

import "os"
import "github.com/codegangsta/cli"

func main() {

  app := cli.NewApp()
  app.Name = "HubToken"
  app.Usage = "Manage GitHub personal access tokens"
  app.Action = func(c *cli.Context) {
    //
  }

  // createOpt := flag.String("create", "token-name", "create a personal access token called `token-name`")
  // flag.Parse()
  // fmt.Println("create token: ", *createOpt)

}
