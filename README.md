This is project to learn go language, so code is not very good.

The main idea is to get notifications if event has not happened, for example:
you have a daily backup setup that notifies sleepingpanda (via http get)
every time it runs. If sleepingpanda does not receive get request to the url
in predefined time period - it sends notification to the email.

the main parts of the program will be:
user registration
endpoint configuration
endpoint storage
notifier

will be using gorm and some web framework

users will have:
  * Username
  * Password
  * Email to send password reset

notifycalls will have:
  * type (email, sms, http call)
  * type value (mail address, sms number, http url)
  * enabled (bool)

notifiers will have:
  * type (email, sms, url)
  * type value (mail address, sms code, url)
  * time period (minutes, hours, days, weeks, months, years)
  * enabled bool
  * notifycall (where to send notifications)
  * last notification sent
  * notification timer start datetime

notifylog:
  * notifier
  * datetime
