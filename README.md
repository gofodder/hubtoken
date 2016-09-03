# HubToken

This command allows you to manage your GitHub personal access tokens from the
command line.

If you are using GitHub Two Factor Authentication you'll know that
quick password access from the command line is not easy.  You will
usually need to setup an Personal Access Token (OAuth2 Token), or an
SSH key.

For temporary access this is inconvenient and will often require you
to jump into a web browser, authenticate with GitHub, get a 2FA code,
and then grab the token from the github web page.

The purpose of Hubtoken is to remove much of this pain by removing the
web browser from the equation and leaving you in the terminal.

## Usage

Create a GitHub personal access token called `token-name`

```
hubtoken -c token-name

# or

hubtoken --create-token token-name
```

The new token will be printed in the terminal after login & 2FA.

- - -

Delete a GitHub personal access token called `token-name`

```
hubtoken -d token-name

# or

hubtoken --delete-token token-name
```

- - -

List your GitHub personal access tokens

```
hubtoken -l

# or

hubtoken --list
```

### Login / 2FA

For each action you'll be asked authenticate with github, enter your
login, password and a GitHub Two Factor OTP code.
