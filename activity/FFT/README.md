# FFT
This activity provides your flogo application the ability to read a binary signal file and perform Fast Fourier Transform
on the input data.


## Installation

```bash
flogo add activity github.com/anshulsharmas/flogo-contrib/activity/FFT
```

## Schema
Inputs and Outputs:

```json
{
  "inputs":[
     {
      "name": "inputFilePath",
      "type": "string"
    },
     {
      "name": "sampleSize",
      "type": "integer"
    },
     {
      "name": "outputFilePath",
      "type": "string"
    }
  ]
}
```
## Settings
| Setting     | Description    |
|:------------|:---------------|
| inputFilePath   | Full file path for inpit binary file |         
| sampleSize      | Number of samples to read  |
| outputFilePath  | Full file path of FFT output file |

## Configuration


### Flow Configuration
Configure a task in flow to perform Fast Fourier Transform

```json
{
  "id": 3,
  "type": 1,
  "activityType": "FFT",
  "name": "Perform FFT",
  "attributes": [
    {
      "name": "inputFilePath",
      "type": "string",
      "value": "/opt/data/input.bin"
    },
    {
      "name": "sampleSize",
      "type": "number",
      "value": 10000
    },
    {
      "name": "outputFilePath",
      "type": "string",
      "value": "/opt/data/output.csv"
    }
  ]
}
```
