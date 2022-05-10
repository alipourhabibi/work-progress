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
$ wp chart draw -name name -time time(regext) -c bar
```
## screenshots
![getting datas](https://github.com/alipourhabibi/simple-blog/blob/master/images/datas.jpg?raw=true)
![bar chart](https://github.com/alipourhabibi/simple-blog/blob/master/images/barchart.jpg?raw=true)

## attention
This app is written using urfave and i made this for practical use and this app is not near being usable so feel free to PR and develop this app.
and for making of this app i used some other opensource modes and i was like being on the shoulder giants.

## contribution
This yet only support bar charts. feel free to PR and add other kind of charts.

## Licence 
MIT

