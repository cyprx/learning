# Simple TCP server-client communication

## Command
- Get random numbers: `05GET04`
    + 05: number of bytes sent
    + GET: command
    + 04: number of random numbers expected

- Response: 3810 20 34 33
    + first two bytes to present number of bytes sent
    + the rests are random numbers requested by client
