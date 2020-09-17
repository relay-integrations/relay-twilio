# twilio-step-send-sms

Send a SMS message using Twilio. 

## Example

```yaml
- name: notify-via-twilio
  image: relaysh/twilio-step-send-sms
  spec:
    twilio: &twilio
      accountSID: !Secret twilioAccountSID
      authToken: !Secret twilioAuthToken
    from: !Secret twilioPhoneNumber
    to: !Parameter phoneNumber
    body: "hello world!"
```
