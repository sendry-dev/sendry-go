# Changelog

All notable changes to `sendry-go` are documented in this file.

The format follows [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [v0.1.0] - 2026-05-11

### Added

- Initial public release.
- `sendry.Client` with resources: Emails, Domains, Templates, APIKeys,
  Webhooks, Analytics, Suppression, Unsubscribes, Contacts, Audiences,
  Campaigns, Billing, Team.
- `VerifyWebhookSignature` helper (HMAC-SHA256, constant-time compare).
- Functional options: `WithBaseURL`, `WithHTTPClient`, `WithTimeout`,
  `WithMaxRetries`.
- Automatic retries on 5xx / network errors with exponential backoff,
  Retry-After honoured on 429.
- Strongly typed error hierarchy: `APIError`, `AuthenticationError`,
  `ValidationError`, `RateLimitError`, `NotFoundError`, `NetworkError`.
