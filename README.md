# HubToken

**NOTE: HubToken is (currently) intended for use only with Github
accounts that have Two Factor Authentication enabled.**

This command allows you to manage your GitHub personal access tokens
from the command line.

If you are using GitHub Two Factor Authentication you'll know that
quick password access from the command line is a bit of a pain.

Usually you'd setup an SSH key, or a Personall access token.  For
temporary access this is a little inconvenient, you'll need to jump
into a web browser, authenticate with GitHub, get a 2FA code, and then
grab the token from the github web page.

Hubtoken removes this friction by keeping you in the terminal, letting
you manage your personal access tokens with minimal fuss.

## Installation

#### Homebrew

```
brew tap gofodder/hubtoken
brew install hubtoken
```

#### Using go

If you have go already installed:

```
go get github.com/gofodder/go-hubtoken
```

## Usage

View the `hubtoken` help & usage page using:

```
hubtoken
```

or help for a specific command

```
hubtoken help [command]
```

Create a GitHub personal access token called `token-name`

```
hubtoken create token-name

# or

hubtoken c token-name
```

The new token will be printed in the terminal after login & 2FA.

- - -

Delete a GitHub personal access token called `token-name`

```
hubtoken delete token-name

# or

hubtoken d token-name
```

- - -

List your GitHub personal access tokens

```
hubtoken list

# or

hubtoken l
```

### Login / 2FA

For each action you'll be asked authenticate with github, enter your
login, password and a GitHub Two Factor OTP code.

### Development task list

- [ ] Setup testing
    - [ ] Mock GitHub endpoints

        NOTE: Probably use the same/similar testing setup as go-github (TBD)

        - [ ] Authentication
        - [ ] Authorizations
            - [ ] Create
            - [ ] Delete
            - [ ] List

    - [ ] Add regression tests for
        - [ ] Login
        - [ ] Create token
        - [ ] Delete token
        - [ ] List tokens

- [ ] Test drive new dev work
- [x] Create token NAME mandatory / prompt if missing. (TBD)
- [x] Delete token NAME mandatory / prompt if missing. (TBD)
- [ ] Input validation
    - [ ] don't allow certain chars (check github for token name rules)
    - [ ] check that duplicate token names cannot be used
- [-] Better error messages

#### 1.0.1

- [x] Pretty ansi color messages

#### 0.1.0

- [x] Create Token
    - ~~go get octokit golang~~
    - [x] go get go-github
    - ~~Use `flag` for opt parsing (import "flag")~~ use `cli` instead
    - [x] Use `codegangsta/cli` as a cli front end
    - [x] Get login from terminal
    - [x] Get password from terminal and turn off echo
    - [x] Get otp/2fa code from terminal
    - ~~Connect to github via octokit~~ use go-github instead
    - [x] connect to github via go-github using OTP/2FA
    - [x] Create/retrieve token
    - [x] Print token to STDOUT
- [x] delete token
- [x] list tokens

[Go]:https://golang.org/
