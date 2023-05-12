## How to run
### Run all services and graphql playground
- System: linux
- Command:
```
make up
```

### Run single services
- Supported profile: flight, user, customer, graphql, booking
- Example for user service:
```
make upProfile Profile="user"
```

## Planned feature
- Implement transaction.
- Add more fields validation.