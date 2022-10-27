# Plan

## Architecture
- Tests
- Deploy as lambda with serverless
- Cache results in cloudfront
- Clear cache when blog is uploaded or modified.
    - S3 trigger
    - DynamoDB trigger
- Support multiple blog sources
    - Local file
    - S3
    - DynamoDB
- Log to cloudwatch

## Frontend
- esbuild
