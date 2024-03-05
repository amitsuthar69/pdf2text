### This is a Go web service which converts pdf to text.

To test/use this service, issue an HTTP **POST** request to : 

```
https://pdf2text-a00o.onrender.com/convert
```

with a valid pdf file in the request body as **form-data**.

---

Example usage in POSTMAN:

![image](https://github.com/amitsuthar69/pdf2text/assets/111864432/c533ba6e-2408-47b6-a43f-ba1599c9c1f3)

Note: As this service is deployed on Render, the initial request can take upto 30 secs to respond. So retry if fails in the first attempt.