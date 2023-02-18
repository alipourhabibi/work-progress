# work Progress
This small app lets you add a job, its description the record of your work for that day ( amount of time you did that job)
and then you can use charts to see your performance.

## Usage
Please use -h or --help to see how to use this app.
But for a brief:)
```
$ wp job add -n name -d description -a 12
```

```
$ wp job get -n name -t time(regex)
```
```
$ wp draw -name name
```
```
$ wp draw -t 09-12-2022
```
## screenshots
![table](https://github.com/alipourhabibi/work-progress/blob/master/images/1.jpg?raw=true)
![raw_data](https://github.com/alipourhabibi/work-progress/blob/master/images/2.jpg?raw=true)
## Using termgraph
```
$ wp draw -name name | termgraph
```
![termgraph](https://github.com/alipourhabibi/work-progress/blob/master/images/3.jpg?raw=true)

## attention
This app is written using urfave and i made this for practical use and this app is not near being usable so feel free to PR and develop this app.
