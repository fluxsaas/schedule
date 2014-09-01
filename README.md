## Sample Go Project

This library should be able to build a schedule from workers and shifts.

This first version implements following simple logic:

1. collect all day ranges in schedule (aka `shifts`)
2. iterate over each shift and assign the first available workers.

Enhanced version should include more options:
* Departments
* WorkingHours
* Rotating
* Preferences
* Blocking days
* ...

Maybe there is a more generic approach on restriction (`rule sets`) ?

We should be able to provide input as JSON:

`scheduler/fixtures/schedule.json`

`````json
{
    [
      "starts_at" "2014-08-25T04:30:00+02:00",
      "ends_at" "2014-08-25T08:30:00+02:00"
    ]
}
`````

`scheduler/fixtures/workers.json`

`````json
{
    "id": 1,
    "name": "kalle saas",
    "available": [
        {
            "wday": 0,
            "starts_at": 800,
            "ends_at": 1630
        }
    ]
}
`````
