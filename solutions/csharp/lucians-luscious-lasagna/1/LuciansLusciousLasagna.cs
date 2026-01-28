class Lasagna
{
    public int ExpectedMinutesInOven() => 40;

    public int RemainingMinutesInOven(int passedTime) => ExpectedMinutesInOven() - passedTime;

    public int PreparationTimeInMinutes(int layerCount) => layerCount * 2;

    public int ElapsedTimeInMinutes(int layerCount, int passedTime) => PreparationTimeInMinutes(layerCount) + passedTime;
}
