# Use a slim Go image to build the application
FROM golang:1.22-alpine AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Build the Go app, creating a static binary
RUN go build -o /app/srenity-mvp-backend main.go

# Use a minimal, non-root image for the final stage
FROM gcr.io/distroless/static-debian11
WORKDIR /
COPY --from=build /app/srenity-mvp-backend .
EXPOSE 8000
CMD ["/srenity-mvp-backend"]