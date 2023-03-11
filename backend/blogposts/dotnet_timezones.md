We've decided at work to only store UTC-times at work.
So naturally we need to convert the times to the viewers time.

At one point I came a cross a situation where I needed to convert from UTC, to a certain timezone, and then back.
Seems easy enough.

I did not think much of it but let intellisense help me find what I needed, and so, I ended up with the following:

````
TimeZoneInfo tz = TimeZoneInfo.FindSystemTimeZoneById("W. Europe Standard Time");
var localNow = TimeZoneInfo.ConvertTimeFromUtc(now, tz);
var nowConvertedBack = localNow.ToUniversalTime();
````

**Can you spot the error?**

---


`localNow` is a DateTime, and such, does not contain any timezone info.
Then why does it have a function `ToUniversalTime`?
Well I think that it shouldn't as it can lead to unexpected result.
For a detailed explanation see: https://learn.microsoft.com/en-us/dotnet/api/system.timezoneinfo.converttimefromutc?view=net-7.0#remarks

Anyhow, the solution is simply to not rely on intellisense and end up with ToUniversalTime() on a DateTime but rather use the `TimezoneInfo` we already have:
`var nowConvertedBack = TimeZoneInfo.ConvertTimeToUtc(localNow, tz);`

Simple enough.

---

Invalid code:
````
public class Program
{
	public static void Main()
	{
		var now = DateTime.UtcNow;
		TimeZoneInfo tz = TimeZoneInfo.FindSystemTimeZoneById("W. Europe Standard Time");
		var localNow = TimeZoneInfo.ConvertTimeFromUtc(now, tz);
		var nowConvertedBack = localNow.ToUniversalTime();
		
		Console.WriteLine("UTC now: " + now);
		Console.WriteLine("local now: " + localNow);
		Console.WriteLine("local now kind: " + localNow.Kind);
		Console.WriteLine("UTC now: " + nowConvertedBack);
	}
}
````

Correct code:
`````
public static void Main()
	{
		var now = DateTime.UtcNow;
		TimeZoneInfo tz = TimeZoneInfo.FindSystemTimeZoneById("W. Europe Standard Time");
		var localNow = TimeZoneInfo.ConvertTimeFromUtc(now, tz);
		var nowConvertedBack = TimeZoneInfo.ConvertTimeToUtc(localNow, tz);
		
		Console.WriteLine("UTC now: " + now);
		Console.WriteLine("local now: " + localNow);
		Console.WriteLine("UTC now: " + nowConvertedBack);
	}
`````
