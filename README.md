# Notes DevOpsDays Zürich 2022

## The Socio-Technical Path to High-Performing Teams Charity Majors
The happiness and satisfaction of your users and your engineers tend to rise
and fall in tandem.

6570 times faster lead time from commit to deploy.

Only hire from the best schools - stupid

Your productivity tends to rise or fall to match that of the team you join.

Three elements of the sociotechnical system

 - Your Team
 - Your Tools
 - Production

=> Feedback Loops

CI is only the appetizer.

CD as in Continuous Deployment!

15 minutes to prod is nice, but it's more important that it's predictable.

Undeployed software ages like fine milk.

[How much is your fear of continuous deployment costing you](https://charity.wtf/2021/02/19/how-much-is-your-fear-costing-you/)

## GitOps for the People Lian Li
Never speak to management when addressing delivery pains:

- Quality Assurance (testers manually deploy on the test environment)
- Diverting environments
- Compliance

Suggestions:

- QA -> Automated Testing, also, embed them
- Diverting environments -> Kubernetes
- Compliance -> GitOps

Manual process: they were first looking at stuff that noone really wanted to do.

The ops team only has to approve the pull request, then it goes to prod.

JiraOps! Ticket-driven delivery.

## Hacking Kubernetes - Live Demo Marathon Andrew Martin from controlplane
[@captainhashjack](https://twitter.com/captainhashjack)

What are your adversaries - script kiddies to nation states

Supply chain attack - NPM modules

### package.json has pre-install hooks

How to fix:
- 2FA for signing keys and SSH
- no plaintext tokens

Docker images: typo-squatting

### Dirty pipe
replacing runc

[Canary Tokens](https://www.canarytokens.org/generate)

## Cybersecurity Beyond Automation - David Sommer and Derek Yu
Don't forget the legacy systems, as long as they're active people have the
easiest way in there.

Just having the database backed up is not really a good strategy, you have
to be able to bootstrap the whole system.

security.txt

https://www.google.com/.well-known/security.txt

-> we should really add that to our product

SBOM -> Software Bill Of Materials, that way at least you can filter on CVEs linked on that.

## DevOps culture at Amazon - Ramon Lopez Narvaez
Werner Vogels: You build it, you run it. https://queue.acm.org/detail.cfm?id=1142065

Not only APIs as interface, but also: office hours, ticketing system etc.

They will *never* have to talk to an Ops team.

### Operation Excellence Reviews
Usually once a week.

### Deployment Strategies
Multi-staged: Problem is your beta is also another team's dependency.

https://aws.amazon.com/architecture/well-architected/?nc1=h_ls&wa-lens-whitepapers.sort-by=item.additionalFields.sortDate&wa-lens-whitepapers.sort-order=desc

## Lightning Talks

### From IT Change Management to Change Enablement
Developers *can* deploy to prod.

Score card to trust application teams.

### Gamechanger skills for the future Tobias Hauk
Where's the humanity.

Coping with complexity

Mindfulness

"Busy is the new stupid"

### Are containers and K8s overhyped? Sebastian Mangelkramer
Overhyped? Sure. You need a reason.

### Mentorship programs Oleg Nenashev
User => Committer => Maintainer => Leader

Mentorship programs are great for the community.

### Big goal requires big leadership Achim Töper

### Power Coders
-> Christina Gräni


## From the Sponsor Booth
CheckPoint smart workflows and API endpoints for policy changes.

## One evolution of the Network Engineer - Jose Leitao

### Network knowledge
protocols, topologies, how the internet works and you

### How to interact with devices

## SBB BizDevOps
650 million budget for IT

## Ignites
### T-Shaped people. Modern vision or history repeats itself? Andriy Bilous
https://en.wikipedia.org/wiki/TRIZ


