Assumptions:
A circle has 360 degrees, the objects are in a circular orbit arround the sun therefore Length of track is 360. 
Year length is 365. 
Objects on same degrees are on same position. 
Everything is BASED on a cycle approach.
every event happens n time on a cycle, therefore multiplying n * amountCycles = finalAmount


drought season:
//DONE have to correct the implementation so that it checks for every position with each other, as to find the one with least intersections.
//DONE have to calculate for N cycles. aka years

Rain season:

Most Rain:

Optimum climate:
If there's time, improve for better checking. As of now, since the checks are being done once a day, it's a tad imprecise. the refactor would be to be able to check multiple times a day. Not terrible but an improvement.
// DONE 
this is when the three points are collinear. AND! they are NOT collinear with the sun. (test for the remaining point? ORRR use the days there's drough! on those days, we should NOT check for collinearity. Maybe implement both approaches, since one's more efficient and the other one is a compromise.)
//DONE First have to convert the polar coordinates into cartesian coordinates. 

//DONE Current problems: can't execute a dinamic function (which depends on wich event I'm trying to check for) and at the same time save on a "global" var.
-------


problems: 
could take less than a day to complete a period.
No validations anywhere
The program accepts only one sun on the system, and on some locations, only three planets. 
what happens if it takes less than a cycle to complete the years given? 
