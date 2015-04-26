var theater = new TheaterJS();

//set actors
theater
    .describe("ThingsLab", 1, "#typo")
    .describe("InnocentUser", 1, "#corrector")
    .describe("y-sir", .6 , "#yes");

//scenario
theater
    .write("ThingsLab:It looks like you made a tpyo.", 600)
    .write("InnocentUser:You mean,", 600, " \"a typo\"?")
    .write("y-sir:", 200, "Yes...");