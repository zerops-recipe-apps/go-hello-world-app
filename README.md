<!-- #ZEROPS_REMOVE_START# -->
# Go Hello World Recipe App
Simple Go API with single endpoint that reads from and writes to a PostgreSQL database. Used within [Go Hello World recipe](https://app.zerops.io/recipes/go-hello-world) for [Zerops](https://zerops.io) platform.

⬇️ **Full recipe page and deploy with one-click**

[![Deploy on Zerops](https://github.com/zeropsio/recipe-shared-assets/blob/main/deploy-button/light/deploy-button.svg)](https://app.zerops.io/recipes/go-hello-world?environment=small-production)

![nestjs](https://github.com/zeropsio/recipe-shared-assets/blob/main/covers/svg/cover-go.svg)

## Integration Guide
<!-- #ZEROPS_REMOVE_END# -->

> [!TIP]
> If you've deployed the recipe with one-click, it used [this repository](https://github.com/zerops-recipe-apps/go-hello-world-app) to deploy the app from. You can either use this repository as a template, or follow the guide on how to integrate similar setup to Zerops. If you want to more advanced examples, see all [Go recipes](https://app.zerops.io/recipes?lf=go) on Zerops.

### 1. Adding `zerops.yaml`
The main application configuration file you place at the root of your repository, it tells Zerops how to build, deploy and run your application.

```yaml
zerops:
  # Defining production setup, that will run the built application.
  - setup: prod
    build:
      # Using Go build base image, that has Go (with build tools) pre-installed.
      base: go@1
      buildCommands:
        # So we can just simply build the app using the 'go' command.
        - go build -o app main.go
      # All we need to deploy to runtime containers is the built 'app' binary.
      # Package only it.
      deployFiles: ./app
    run:
      # Now, we have to say into which base image we want to deploy our app.
      # Since Go is a compiled language that produces a binary,
      # we can grab a lightweight Linux distribution to run the binary in.
      base: alpine@3.21
      # Defining ports that can be accessed from outside the application container.
      ports:
        - port: 8080
          # Our app is an HTTP API. Mark the port as HTTP
          # so we can possibly enable public HTTPS access.
          httpSupport: true
      # Adding environment variables.
      # Note that we reference database service environment variables,
      # that are automatically generated and accessible for all PostgreSQL services.
      envVariables:
        DB_NAME: db
        DB_HOST: ${db_hostname}
        # For example, this 'DB_PORT' env will resolve to 6543 in case of a PostgreSQL database.
        DB_PORT: ${db_port}
        DB_USER: ${db_user}
        DB_PASS: ${db_password}
      # This is how we execute our app process.
      # We build the 'app' artifact above.
      start: ./app
  
  # Dev setup is for remote development or AI agent use-cases.
  - setup: dev
    build:
      base: go@1
      # Start by packaging all the application source code
      # in the repository, so we can work on it inside the runtime container.
      # No build steps are needed, since we only care about source code.
      deployFiles: .
    run:
      base: go@1
      # We would also like to test and try the app from the outside,
      # make the development port accessible.
      ports:
        - port: 8080
          httpSupport: true
      # Use the same environment variables for development,
      # they will be available in the environment of spawned shells, IDEs or AI agents.
      envVariables:
        DB_NAME: db
        DB_HOST: ${db_hostname}
        DB_PORT: ${db_port}
        DB_USER: ${db_user}
        DB_PASS: ${db_password}
      # We don't want to run anything - we will execute our
      # build, test and run commands manually inside the container.
      # Start command will be optional in the future. Use noop dummy command.
      start: zsc noop --silent
```
