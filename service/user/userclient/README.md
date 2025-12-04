# User Client Package

This directory contains the generated gRPC client code for the User service.

## Current Status

Currently, this directory contains a **placeholder file** (`userclient.go`) that allows the code to compile before generating the actual code from the proto file.

## Generating the Actual Code

To generate the real client code from the proto file, run:

```bash
./scripts/generate-service.sh user
```

Or manually:

```bash
goctl rpc protoc service/user/user.proto \
  --go_out=service/user \
  --go-grpc_out=service/user \
  --zrpc_out=service/user \
  --style gozero
```

## After Generation

After running the generation command:
- The placeholder `userclient.go` will be replaced with generated code
- Multiple files will be created (`.pb.go`, `_grpc.pb.go`, etc.)
- The code will have full gRPC functionality

## Note

The placeholder file is intentionally simple and does not provide actual functionality. It only allows the codebase to compile. After generation, you'll have the full gRPC client implementation.

