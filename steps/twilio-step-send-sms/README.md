# twilio-step-send-sms

Send a SMS message using Twilio. 

## Specification

| Setting | Child setting | Data type | Description | Default | Required |
|---------|---------------|-----------|-------------|---------|----------|
| `twilio` || mapping | A mapping of Twilio account configuration. | None | True |
|| `accountSID` | string | Twilio account SID. Use the Secrets sidebar to configure as Secret. | None | True |
|| `authToken` | string | Twilio auth token. Use the Secrets sidebar to configure as Secret | None | True |
| `from` || string | Phone number to send from. | None | True |
| `to` || string | Phone number to send to. | None | True |
| `body` || string | Message body | None | True |

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
