# **Gecko Open API \- General Guide**

## **1\. Overview**

Welcome to the Gecko Open API\! This platform is dedicated to providing developers with powerful and flexible interfaces to seamlessly integrate your applications with Gecko's core functionalities.

**Core Capabilities**:

* Publish posts with rich media content via the **HTTP API**.

Currently, we offer the following Community(Forum) related APIs:

* Forum\_post\_createPost: To create a new post.  
* Forum\_upload\_attachment: To upload image or video attachments for a post.

## **2\. Quick Start (Get Started in 5 Minutes)**

This section will guide you through the process of making your first API call. For detailed interface parameters and examples, please refer to the specific API documentation.

### **Step 1: Obtain Your API Credentials**

Currently, the API Key and Secret Key are created and provided to developers by a **Gecko Administrator**. Please contact our administrators to obtain your credentials.

**Please Note**: The feature to self-apply for an API Key within the application is a planned feature and is currently under active development.

Once you have your credentials, be sure to store the Secret Key securely and never expose it in client-side code or public channels.

### **Step 2: Prepare the Request Data**

Prepare the corresponding request body (for POST/PUT requests) or query parameters (for GET requests) according to the specific API you are calling.

* For example, the request body for creating a post might include fields like contents and tags.

### **Step 3: Generate the Signature**

The Gecko API signature is generated using the HMAC-SHA256 algorithm to ensure the integrity and authenticity of the request.

1. **Construct the String to Sign**: Concatenate the following four parts in order, separated by a colon (:):  
   * The value of the Timestamp HTTP Header (as a string, i.e., Unix timestamp in seconds).  
   * The HTTP Method (e.g., POST, in uppercase).  
   * The Request URI (the path \+ query string from the URL, e.g., "/openapi/forum/post/createPost" from "[https://xxx.com/openapi/forum/post/createPost](https://www.google.com/search?q=https://xxx.com/openapi/forum/post/createPost)").  
   * The Request Body (the raw JSON string of the request. This should be an empty string "" if the request's **ContentType is multipart/form-data**).

   **Important Notes**:

   * The Request Body must be the original JSON string, not a parsed object. We have chosen not to include multipart/form-data in the signature calculation due to its potential size.  
   * The Request URI does **not** include the host address.  
2. **Calculate the MD5 Hash**: Compute the MD5 hash of the constructed string from the previous step and get its hexadecimal representation.  
3. **Calculate the HMAC-SHA256 Signature**: Use your Secret Key to encrypt the MD5 hexadecimal string from step 2 using HMAC-SHA256. The final result should be represented as a hexadecimal string.

### **Step 4: Send the HTTP Request**

Send the request using cURL or any HTTP client. Be sure to include the Api-Key, Signature, and Timestamp in the HTTP headers.

\# Example: General HTTP POST Request  
curl \-X POST 'https://api.example.com/openapi/v1/some\_endpoint' \\  
\-H 'Content-Type: application/json' \\  
\-H 'Api-Key: YOUR\_API\_KEY' \\  
\-H 'Signature: GENERATED\_SIGNATURE' \\  
\-H 'Timestamp: CURRENT\_UNIX\_TIMESTAMP\_SECONDS' \\  
\-d 'YOUR\_REQUEST\_BODY\_JSON'

* Replace YOUR\_API\_KEY, GENERATED\_SIGNATURE, CURRENT\_UNIX\_TIMESTAMP\_SECONDS, and YOUR\_REQUEST\_BODY\_JSON with your actual values.

### **Step 5: Review the Result**

If the request is successful, you will typically receive a response in the following format:

{  
  "code": 0,  
  "msg": "Success",  
  "data": {  
    "someId": "value\_of\_id"  
  }  
}

For specific response content and fields, please refer to the detailed documentation for each API.

## **3\. Authentication Mechanism**

The Gecko API uses an **API Key \+ Secret Key** signature method for authentication to ensure communication security and data integrity.

* **Api-Key**: Your unique API credential, passed via the Api-Key HTTP header to identify your application.  
* **Secret Key**: The key paired with your Api-Key, used to generate the request signature. **It must never be exposed or hardcoded in any client-side code or public channels**.  
* **Timestamp**: The Unix timestamp (in seconds) when the request was initiated, passed via the Timestamp HTTP header. The server will reject requests with a timestamp that differs from the current server time by more than **5 minutes** to prevent replay attacks.  
* **Signature**: An HMAC-SHA256 signature generated based on the Secret Key and specific request parameters, passed via the Signature HTTP header.

**Signature Generation Logic**:

1. **Collect Data to Sign**: Concatenate the following strings in order:  
   * The value of the Timestamp HTTP Header (as a string).  
   * The HTTP Method (uppercase, e.g., GET, POST).  
   * The Request URI (e.g., /path/to/api?param1=value1\&param2=value2).  
   * The Request Body (the raw content of the request body, usually a JSON string. If the request has no body, this is an empty string "").  
2. **Concatenate the String**: Join the four parts using a colon (:) to form a single string to be signed.  
   * Format: TimestampValue:HTTP\_Method:Request\_URI:Request\_Body\_String  
3. **Calculate MD5**: Perform an MD5 hash on the concatenated string to get its hexadecimal representation.  
4. **Calculate HMAC-SHA256**: Use your Secret Key as the key to perform an HMAC-SHA256 encryption on the MD5 hexadecimal string from step 3\.  
5. **Output**: Convert the result of the HMAC-SHA256 to a hexadecimal string, which is the final Signature.

## **4\. Core Concepts**

### **4.1 Unified ID Format**

For ease of management and scalability, all resource unique identifiers (IDs) on the platform use a prefixed string format:

* **Post ID**: Starts with post\_, e.g., post\_ghi789jkl012

## **5\. Error Codes**

When an API request fails, the Gecko Open API will return a JSON response containing code and msg fields. The code field indicates the specific error type:

| Code | Message | Description |
| :---- | :---- | :---- |
| 10001 | Invalid API Key | The provided API Key does not exist or has been disabled. |
| 10002 | Invalid Signature | The signature is invalid. Please check the signing algorithm and the string-to-sign construction. |
| 10003 | Timestamp Expired | The request timestamp has expired or differs too much from the server time. |
| 10004 | Permission Denied | This API Key does not have permission to perform this action. |
| 20001 | Invalid Parameters | Request parameters are incorrect, e.g., missing required fields, incorrect format, or invalid values. |
| 20002 | Rate Limit Exceeded | The API call frequency limit has been exceeded. Please try again later. |
| 30001 | Resource Not Found | The requested resource does not exist, e.g., an invalid Post ID. |
| 50000 | Internal Server Error | An unknown internal server error occurred. |

## **6\. Code Samples (Signature Function)**

The following code samples demonstrate how to construct the string-to-sign and generate the signature according to the logic described in this document.

### **Node.js / JavaScript**

const crypto \= require('crypto');

/\*\*  
 \* Generates an HMAC-SHA256 signature for Gecko API requests, with an intermediate MD5 step.  
 \* @param {string} timestampStr \- The Unix timestamp (in seconds) as a string, from the 'Timestamp' HTTP header.  
 \* @param {string} method \- The HTTP method (e.g., 'POST'), in uppercase.  
 \* @param {string} requestURI \- The request URI (e.g., '/openapi/forum/post/createPost').   
 \* @param {string} bodyString \- The raw request body as a string. Empty string if content type is multipart/form-data.  
 \* @param {string} secretKey \- Your Secret Key.  
 \* @returns {string} The hex-encoded signature.  
 \*/  
function generateSignature(timestampStr, method, requestURI, bodyString, secretKey) {  
  // Construct the original message string according to the API documentation  
  const originalMessage \= \`${timestampStr}:${method}:${requestURI}:${bodyString}\`;  
    
  // Calculate MD5 hash of the original message  
  const md5Hash \= crypto.createHash('md5').update(originalMessage).digest('hex');

  // Create HMAC-SHA256 hash using the MD5 hash as the input to HMAC  
  return crypto  
    .createHmac('sha256', secretKey)  
    .update(md5Hash)  
    .digest('hex');  
}

### **Python**

import hmac  
import hashlib  
import time  
import json

def generate\_signature(timestamp\_str: str, method: str, request\_uri: str, body\_string: str, secret\_key: str) \-\> str:  
    """  
    Generates an HMAC-SHA256 signature for Gecko API requests, with an intermediate MD5 step.  
    :param timestamp\_str: The Unix timestamp (in seconds) as a string, from the 'Timestamp' HTTP header.  
    :param method: The HTTP method (e.g., 'POST'), in uppercase.  
    :param request\_uri: The request URI (e.g., '/openapi/forum/post/createPost').  
    :param body\_string: The raw request body as a string. Empty string if content type is multipart/form-data.  
    :param secret\_key: Your Secret Key.  
    :return: The hex-encoded signature.  
    """  
    \# Construct the original message string according to the API documentation  
    original\_message \= f"{timestamp\_str}:{method}:{request\_uri}:{body\_string}"  
      
    \# Calculate MD5 hash of the original message  
    md5\_hash \= hashlib.md5(original\_message.encode('utf-8')).hexdigest()

    \# Create HMAC-SHA256 hash using the MD5 hash as the input to HMAC  
    return hmac.new(  
        secret\_key.encode('utf-8'),  
        md5\_hash.encode('utf-8'), \# Use the MD5 hash as the input for HMAC  
        hashlib.sha256  
    ).hexdigest()

## **7\. Rate Limiting**

To ensure service quality, we impose limits on API call frequency. By default, each API Key is limited to **1200** calls per minute. Please monitor the X-RateLimit-\* headers in the HTTP response to understand your current limits and usage.

## **8\. Security Best Practices**

* **Keep Your Secret Key Confidential**: The Secret Key grants full operational access to your account. **Never** hardcode it into any front-end code or commit it to public repositories (like GitHub). Ensure that signing operations are performed only on your secure back-end server.  
* **Back-End Signing**: The signature generation process should always be completed on your trusted back-end server. Avoid performing signing on the client-side (e.g., in a browser or mobile app) to prevent Secret Key leakage.  
* **Use an IP Whitelist**: To enhance security, we strongly recommend configuring an IP whitelist for your Key in the Gecko API management page. This will restrict API access to requests originating only from your specified server IP addresses.  
* **Rotate Keys Periodically**: Regularly generate new Api-Key and Secret Key pairs and deactivate old ones. This helps mitigate the risk of a compromised key, as its effective lifetime is limited even if accidentally exposed.  
* **Time Synchronization**: Ensure your server's time is synchronized with the Gecko API server time. A significant time skew (more than 5 minutes) will cause signature validation to fail, resulting in your requests being rejected.
