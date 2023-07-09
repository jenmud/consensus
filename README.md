### Consensus is still under development

# Consensus

Consensus helps building and maintaining software by describing the expected behavior.
It focuses on the high level expectations rather then low level details encouraging a wider
participation from stakeholders, collaborators, and developers.

It aims to provide a platform for creating a project as an `Epic` and instead of tickets/stories/tasks
as the traditional tracking systems do, it focuses on the expected behavior using the [Gherkin language](https://en.wikipedia.org/wiki/Cucumber_(software)#Gherkin_language). Ticket/stories/tasks become [Features](https://en.wikipedia.org/wiki/Cucumber_(software)#Features), [Scenarios](https://en.wikipedia.org/wiki/Cucumber_(software)#Scenarios) and [Steps](https://en.wikipedia.org/wiki/Cucumber_(software)#Steps).

Each *Feature*,  *Scenario*, or *Step* should be small and focuses on one thing and encourages many smaller *Features*, *Scenarios* or *Steps*.

Developers will then pull down the *Features*, *Scenarios* or *Steps* via the API and they will code their applications against the expected behavior and uploading the test results. These results can be used for indicating progress and completion.
