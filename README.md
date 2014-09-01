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
