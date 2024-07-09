# Moonphase

## Build
```
go build
```

## Usage
### Found the phase for date
````
./moonphase -a phase -d "2024-06-15 12:15:30"        
The phase is first quarter for date 2024-06-15
````

### Found the next date of phase after date
````
./moonphase -a date -d "2024-06-15 12:15:30" -p "full moon"
The next full moon is for date 2024-06-18
````
