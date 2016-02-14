# workbench

`workbench` is a command line tool for quickly launching AWS instances,
attaching to tmux sessions, creating snapshots, and more.

Often, my daily data analysis workflow will consist of launching an AWS
instance in the morning and connecting to a running `tmux` session repeatedly
througout the day to run and improve my data analysis pipeline. Towards the
end of the day, I create a snapshot of my running instance and terminate it.
The snapshot is usually the basis for the instance I launch the next day.

I wrote `workbench` to streamline this workflow, so that it can be performed
in seconds and without recourse to the AWS web interface.

You can install it with
```bash
go get github.com/ibab/workbench
```

Before using `workbench`, make sure that you have stored your AWS credentials in `~/.aws/credentials`:
```
[default]
aws_access_key_id = MYAWSACCESSKEY
aws_secret_access_key = MYAWSSECRETKEY
```

In order to list all currently running instances, execute
```bash
$ workbench status
## instances: 1
[1] R ec2-12-34-56-789.compute-1.amazonaws.com
```

To see all available personal AWS images, run
```bash
$ workbench images
## images: 1
[1] ubuntu-15.10-cuda-dev
```

In order to open a new tmux session or attach to an existing one, run
```bash
$ workbench attach 1
```
with the index of the instance you want to attach to.

In order to launch a new instance (currently only spot requests are performed), run
```bash
$ workbench launch 1
```
with the index of the image you want to base the instance on.
`workbench` will then ask you for the instance type you want to use and allow you to set
the availability zone and spot price.
The spot request is currently limited to an arbitary 10h.

If you want to terminate an instance, run
```bash
$ workbench terminate 1
```
with the index of the instance you want to terminate.
This will also cancel the spot request.

## Commands that still need to be implemented:
```
$ workbench snap
$ workbench rmi
```

