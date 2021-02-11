# logrus-coralogix-cpu
A reproduction of High CPU utilization with logrus and Coralogix Hook

During this test - make sure to run `htop` command or similar to see the CPU utilization.

1. run:
```sh
make build
```

2. On one terminal, run:
```sh
make run
```

3. On another terminal, run:
```sh
make profile
```

Observe the CPU.

4. Uncomment line 63 in main.go, and repeat steps 1-3
