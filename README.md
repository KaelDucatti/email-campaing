# Email Campaign Service

A high-performance email campaign management system built with Go, featuring authentication via Keycloak and comprehensive observability.

## Architecture Overview

The system consists of the following components:

- **User Interface**: Frontend application for campaign management
- **E-mail Campaign Service**: Core Go backend service handling campaign logic
- **Keycloak**: Authentication and authorization server
- **Database**: Persistent storage for campaigns and user data
- **E-mail Provider**: Integration for sending emails
- **Observability Worker**: Monitoring and metrics collection

## Tech Stack

- **Backend**: Go 1.21+
- **HTTP Router**: Chi
- **Authentication**: Keycloak
- **Database**: PostgreSQL (recommended)
- **Observability**: Prometheus, Grafana (or your preferred stack)

## Features

- 🔐 Secure authentication with Keycloak
- 📧 Email campaign creation and management
- 📊 Real-time campaign analytics
- 🔍 Comprehensive observability and monitoring
- 🚀 High-performance API built with Go and Chi
- 📝 Database persistence for campaigns and templates

## Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose
- Keycloak instance (or use the provided Docker setup)
- PostgreSQL database
- SMTP server or email service provider credentials

## Getting Started

### 1. Clone the repository

```bash
git clone <repository-url>
cd email-campaign
```

### 2. Environment Setup

Copy the example environment file and configure your settings:

```bash
cp .env.example .env
```

Edit `.env` with your configuration:

```env
# Server Configuration
PORT=8080
ENVIRONMENT=development

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=emailcampaign
DB_PASSWORD=your_password
DB_NAME=emailcampaign

# Keycloak
KEYCLOAK_URL=http://localhost:8081
KEYCLOAK_REALM=email-campaign
KEYCLOAK_CLIENT_ID=campaign-service
KEYCLOAK_CLIENT_SECRET=your_client_secret

# Email Provider
SMTP_HOST=smtp.example.com
SMTP_PORT=587
SMTP_USER=your_email@example.com
SMTP_PASSWORD=your_smtp_password

# Observability
METRICS_PORT=9090
TELEMETRY_ENDPOINT=http://localhost:4318
```

### 3. Start Infrastructure Services

```bash
docker-compose up -d
```

This will start:
- PostgreSQL database
- Keycloak server
- Prometheus
- Grafana

### 4. Run Database Migrations

```bash
make migrate-up
```

### 5. Start the Application

```bash
make run
```

Or build and run:

```bash
make build
./bin/email-campaign
```

## API Documentation

### Authentication

All API endpoints require authentication via Keycloak JWT tokens.

Include the token in the Authorization header:
```
Authorization: Bearer <your_jwt_token>
```

### Endpoints

#### Campaigns

- `GET /api/v1/campaigns` - List all campaigns
- `POST /api/v1/campaigns` - Create a new campaign
- `GET /api/v1/campaigns/{id}` - Get campaign details
- `PUT /api/v1/campaigns/{id}` - Update a campaign
- `DELETE /api/v1/campaigns/{id}` - Delete a campaign
- `POST /api/v1/campaigns/{id}/send` - Send campaign emails

#### Templates

- `GET /api/v1/templates` - List email templates
- `POST /api/v1/templates` - Create a new template
- `GET /api/v1/templates/{id}` - Get template details
- `PUT /api/v1/templates/{id}` - Update a template
- `DELETE /api/v1/templates/{id}` - Delete a template

#### Analytics

- `GET /api/v1/campaigns/{id}/analytics` - Get campaign analytics
- `GET /api/v1/campaigns/{id}/metrics` - Get detailed metrics

## Development

### Project Structure

```
.
├── cmd/
│   └── server/           # Application entrypoint
├── internal/
│   ├── api/              # HTTP handlers and routes
│   ├── auth/             # Keycloak integration
│   ├── campaign/         # Campaign business logic
│   ├── database/         # Database models and queries
│   ├── email/            # Email provider integration
│   └── observability/    # Metrics and tracing
├── migrations/           # Database migrations
├── pkg/                  # Shared packages
├── docker-compose.yml
├── Makefile
└── README.md
```

### Running Tests

```bash
make test
```

Run tests with coverage:

```bash
make test-coverage
```

### Code Quality

Run linting:

```bash
make lint
```

Format code:

```bash
make fmt
```

## Monitoring and Observability

### Metrics

The application exposes Prometheus metrics at `/metrics` on the metrics port (default: 9090).

Key metrics:
- Campaign creation/send rates
- Email delivery success/failure rates
- API response times
- Database query performance

### Grafana Dashboards

Access Grafana at `http://localhost:3000` (default credentials: admin/admin)

Pre-configured dashboards include:
- Campaign Performance
- System Health
- API Metrics
- Database Performance

## Deployment

### Docker

Build the Docker image:

```bash
docker build -t email-campaign:latest .
```

Run with Docker:

```bash
docker run -p 8080:8080 --env-file .env email-campaign:latest
```

### Kubernetes

Apply the Kubernetes manifests:

```bash
kubectl apply -f k8s/
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License.

## Support

For support, email support@example.com or open an issue in the repository.