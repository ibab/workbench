# workbench

`workbench` is a command line tool for quickly launching AWS instances,
attaching to tmux sessions, creating snapshots, and more.

Often, my daily data analysis workflow will consist of launching an AWS
instance in the morning and connecting to a running `tmux` session repeatedly
througout the day to run and improve my data analysis pipeline. Towards the
end of the day, I create a snapshot of my running instance and terminate it.
The snapshot is usually the basis for the instance I launch the next day.

I wrote `workbench` to streamline this workflow, so that it can be performed
in seconds and without recurse to the AWS web interface.

You can install it with
```bash
go get github.com/ibab/workbench
```

In order to see list all currently running instances, run
```bash
$ workbench status
## running: 0
No instances are currently running
```

To see all available personal AWS images, run
```bash
$ workbench images
## images: 1
[1] ubuntu-15.10-cuda-dev
```

