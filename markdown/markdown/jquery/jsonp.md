# JSONP

参考：[https://learn.jquery.com/ajax/working-with-jsonp/](https://learn.jquery.com/ajax/working-with-jsonp/)

```javascript
// Using YQL and JSONP
$.ajax({
    url: "http://query.yahooapis.com/v1/public/yql",
 
    // The name of the callback parameter, as specified by the YQL service
    jsonp: "callback",
 
    // Tell jQuery we're expecting JSONP
    dataType: "jsonp",
 
    // Tell YQL what we want and that we want JSON
    data: {
        q: "select title,abstract,url from search.news where query=\"cat\"",
        format: "json"
    },
 
    // Work with the response
    success: function( response ) {
        console.log( response ); // server response
    }
});
```