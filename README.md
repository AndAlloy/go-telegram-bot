# golang-finite-bot
<p>It`s a bot that implements simple FSM with states (aka. StatesGroup() in <a href="https://github.com/aiogram/aiogram">aiogram</a>) 
in <a href="https://github.com/tucnak/telebot">telebot.v3</a></p>
<p>All state configs are in "state.go". In order to control FSM you can:</p>
<ul>
<li>Create new states with function</li>
<li>Enable/disable current state</li>
<li>Save message context</li>
<li>Check if state is active now</li>
</ul>

Docs:
<ul>
<li>Cannot save contex in unactive state</li>
<li>Don`t forget to disable state after context setting</li>
<li>Possible number of states - unlimited</li>
</ul>

Plans:
<ul>
<li>Make a library for simple FSM</li>
<li>Add StatesGroups to connect states in one interface</li>
<li>Add more usefull things to control FSM</li>
</ul>
