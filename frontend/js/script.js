// Example in vanilla JavaScript for sending a message to the backend
async function postMessage(senderId, text) {
    const response = await fetch('/api/messages', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ sender_id: senderId, text: text }),
    });
  
    if (response.ok) {
      const data = await response.json();
      console.log("Message added:", data);
    } else {
      console.error("Failed to add message");
    }
  }

  console.log("Hi!")
  
  // Call this function with the sender ID and message text
  postMessage(1, "Hello, this is a test message!");
  