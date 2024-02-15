### This is a Go web service which converts pdf to text.

To test/use this service, issue an HTTP **POST** request to [https://pdf2text-a00o.onrender.com/convert](https://pdf2text-a00o.onrender.com/convert) with a valid pdf file in the request body as **form-data**.

---

Example usage in POSTMAN:

![image](https://github.com/amitsuthar69/pdf2text/assets/111864432/c533ba6e-2408-47b6-a43f-ba1599c9c1f3)

---

We gonna use this to extract raw text from pdf uploaded by client and then send it back to both client and model, so the client has a preview of his pdf content and model can summarise it. 

The reason why I wrote this service in Go is because idk Python & JavaScript is slow af! 

> Underlying System Design:
> 
> ![image](https://github.com/amitsuthar69/pdf2text/assets/111864432/a8c70fec-8543-466b-9379-3e12bf719778)

