## rand

Learnt about rand function to generate random number. Since, random number generation in computer science isn't truly random, we need to provide a different seed on each `rand` to generate different numbers on each `rand` execution else the same pattern of random numbers will be generated. For that     `time.Now().UnixMill()` which generates the time on millisecond which is unique of course.

## init function

So of course, I already know that the program runs from `main` function at the beginninig. But what was interesting about `init` function was, when the server is started or the code/ package is running the init function gets invoked for the first time only. It kinda works like `useEffect(()=>{},[])` in react. This is helpful when the code is to be initialized for the first time only. Like connecting to database or like in the previous `rand` case seeding the `rand` function. etc.

### fmt.Sprintf

It kinda works like concatenation of string, though i found it kinda requires a placeholder  on the first string to concat the second string.