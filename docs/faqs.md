# Frequently Asked Questions (FAQs)

#### I receive `multiple default rules named allow found`. What happend?

In most cases you have multiple packages with the identical name. Eg.

```
  - id: customer
    raw: |-
      package customer

      default allow = false
  - id: organisations
    raw: |-
      package customer

      default allow = false
```

Each policy needs it unique package name! Otherwise Opa will deny the policy

#### Why the name 'urOpa'?

Uropa means great-grandfather in german. 

