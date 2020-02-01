#Assumptions:
A circle has 360 degrees, the objects are in a circular orbit arround the sun therefore Length of track is 360. 
Year length is 365. 
Objects on same degrees are on same position. 
Everything is BASED on a cycle approach.
every event happens n time on a cycle, therefore multiplying n * amountCycles = finalAmount

#Questions: 
1. ¿Cuántos períodos de sequía habrá?
2. ¿Cuántos períodos de lluvia habrá y qué día será el pico máximo de lluvia?
3. ¿Cuántos períodos de condiciones óptimas de presión y temperatura habrá?

// TODO: Check this! 
All checks should start on day 1 and finish on day 360. Or, start on day 0 and finish on day 359. 

#Objectives:
Four questions // 2 out of 4. 
Job to calculate the days. 
Api access.
Program upload. 

#Objectives++:
Tests.
There are a couple of methods that can go generic, like in the events. 
Concurrency implementation, especially on the job. 
Api explanation file.
try ... catch.. defer.. panic handling 

#Drought season:
//DONE have to correct the implementation so that it checks for every position with each other, as to find the one with least intersections.
//DONE have to calculate for N cycles. aka years

#Rain season:
This one is kind of attached to the most rain day. To calculate this: 
It's rain season when the sun is inside the triangle, so make a method to calculate when the cartesian position of the planets make a triangle with the sun inside.

#Most Rain:
Day where the triangle's perimeter is the biggest. When cycling through the days, check for max perimeter. What if there's more than one day? Maybe at least save the date.

#Optimum climate:
If there's time, improve for better checking. As of now, since the checks are being done once a day, it's a tad imprecise. the refactor would be to be able to check multiple times a day. Not terrible but an improvement.
// DONE 
this is when the three points are collinear. AND! they are NOT collinear with the sun. (test for the remaining point? ORRR use the days there's drough! on those days, we should NOT check for collinearity. Maybe implement both approaches, since one's more efficient and the other one is a compromise.)
//DONE First have to convert the polar coordinates into cartesian coordinates. 

//DONE Current problems: can't execute a dinamic function (which depends on wich event I'm trying to check for) and at the same time save on a "global" var.
--


problems: 
could take less than a day to complete a period.
No validations anywhere
The program accepts only one sun on the system, and on some locations, only three planets. 

// Q: what happens if it takes less than a cycle to complete the years given? 
// R: Maybe check for this at the beginning of the method call, then if cycleDays > days, cycleDays == days (?)

