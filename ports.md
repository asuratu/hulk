# 各个模块的端口划分

### dev port

#### service port

| service name | api service port(1xxx) | rpc service port(2xxx) | other service port(3xxx) | admin service port(4xxx) |
|--------------|------------------------|------------------------|--------------------------|--------------------------|
| app          | 1001                   | -                      | -                        | -                        |
| user         | 1002                   | 2002                   | mq-3002                  | 4002                     |
| order        | 1003                   | 2003                   | mq-3003                  | 4003                     |
| mqueue       | -                      | -                      | job-3002、schedule-3003   | 4004                     |







