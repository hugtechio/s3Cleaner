# S3 Bucket Cleaner
Delete bucket even some objects was remaining

- Deleting all versions
- Deleting all the delete markers
- Deleting all objects
- Deleting Bucket

# Language
golang

# build
```
go build s3Cleaner.go
```

# run
```
go run s3Cleaner.go
```

# Usage
```
s3Cleaner [BUCKET_NAMES_FILE][REGION]
```

example  
```
s3Cleaner buckets.txt us-east-1 
```

- BUCKET_NAMES_FILE - list of bucket name (default: **buckets.txt**)
- REGION - Target region (default: **us-east-1**)