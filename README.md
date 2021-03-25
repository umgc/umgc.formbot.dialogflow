# Team Dialogflow - Form Scriber

This software is free to use by anyone. It comes with no warranties and is provided solely "AS-IS". It may contain significant bugs, or may not even perform the intended tasks, or fail to be fit for any purpose. University of Maryland is not responsible for any shortcomings and the user is solely responsible for the use.

# Agents

A Dialogflow agent is a virtual agent that handles conversations with your end-users.
 
A Dialogflow agent is similar to a human call center agent. You train them both to handle expected conversation scenarios, and your training does not need to be overly explicit.
 
Utterance & Triggers
 

 
 
# Intents


 
 
Intents are setup by telling Dialogflow what intents you want. Dialogflow's engine will handle the intent "matching".
 
An intent categorizes an end-user's intention for one conversation turn. For each agent, you define many intents, where your combined intents can handle a complete conversation. When an end-user writes or says something, referred to as an end-user expression, Dialogflow matches the end-user expression to the best intent in your agent.
 
 
## A basic intent contains the following:

Training phrases: These are example phrases for what end-users might say. When an end-user expression resembles one of these phrases, Dialogflow matches the intent. You don't have to define every possible example, because Dialogflow's built-in machine learning expands on your list with other, similar phrases.
Action: You can define an action for each intent. When an intent is matched, Dialogflow provides the action to your system, and you can use the action to trigger certain actions defined in your system.
Parameters: When an intent is matched at runtime, Dialogflow provides the extracted values from the end-user expression as parameters. Each parameter has a type, called the entity type, which dictates exactly how the data is extracted. Unlike raw end-user input, parameters are structured data that can easily be used to perform some logic or generate responses.
Responses: You define text, speech, or visual responses to return to the end-user. These may provide the end-user with answers, ask the end-user for more information, or terminate the conversation.
 
## A more complex intent may also contain the following:

Contexts: Dialogflow contexts are similar to natural language context. If a person says to you "they are orange", you need context in order to understand what the person is referring to. Similarly, for Dialogflow to handle an end-user expression like that, it needs to be provided with context in order to correctly match an intent.
Events: With events, you can invoke an intent based on something that has happened, instead of what an end-user communicates.
 
# Entities
 

 
 
Entities are essentially variables. Entities can be used to send back hard-coded static responses back the user or sent again to a back-end service to do dynamic processing or other routines.
 
# Contexts
 

 
 
Contexts allow Dialogflow to keep information while iterating through various different intents and holding information as it goes from one intent to the next.
 
 
# Fulfillment
 

 
 
 
This is the last "puzzle piece" of Dialogflow, which allows Dialogflow to connect to your back end systems to create dynamic responses back the user's requests.
 
Fulfillment can work with webhooks and send the user's responses to a custom webservice or google cloud functions.
 

 

 
You can read more on webhoks with example here: https://cloud.google.com/dialogflow/es/docs/fulfillment-webhook
 
 
# Integration - Dialogflow Messenger

 

Documentation: https://cloud.google.com/dialogflow/es/docs/integrations/dialogflow-messenger
 
Google Assistant has been deprecated in favor of Dialogflow Messenger. While it is possible to build your own DF client or use 3rd parties, Google offers several supported clients out of the box.
 
The Dialogflow Messenger integration provides a customizable chat dialog for your agent that can be embedded in your website. The chat dialog is implemented as a dialog window that can be opened and closed by your end-user. When opened, the chat dialog appears above your content in the lower right side of the screen.
