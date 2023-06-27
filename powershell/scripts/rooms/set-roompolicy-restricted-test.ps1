$mail = "test-room-fcada84c-18ae-4c1d-a469-255ea7b99a91@M365x65867376.onmicrosoft.com"
. "$psscriptroot/set-roompolicy-restricted.ps1"  -Mail $mail -Restrictedto alexw,debrab -BookingWindowInDays 601 -MailTip "This is a test room"

