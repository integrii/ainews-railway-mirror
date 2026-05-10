package content

// Section groups related paragraphs under an optional heading.
type Section struct {
	Heading    string
	Paragraphs []string
}

// Post represents a published AINews article.
type Post struct {
	Title     string
	Slug      string
	Date      string
	Tag       string
	Summary   string
	SourceDoc string
	Sections  []Section
}

var posts = []Post{
	{
		Title:     "Apple's Multi-AI Gambit: What iOS 27 Reveals About the Platform Wars",
		Slug:      "apples-multi-ai-gambit-what-ios-27-reveals-about-the-platform-wars",
		Date:      "May 10, 2026",
		Tag:       "Platforms",
		Summary:   "Apple's new Intelligent Services Layer turns Siri into a routing system across Apple, OpenAI, and Anthropic models, setting up a direct architectural contrast with Google's Gemini-first Android strategy.",
		SourceDoc: "https://docs.google.com/document/d/1dEJOQa4GewWrSH8N6sCUEnMkJZWwiy2kOlHw3VcjaGw/edit",
		Sections: []Section{
			{
				Paragraphs: []string{
					"Apple's first iOS 27 beta arrives just ahead of Google's May 12 Android Show and makes one strategic point unmistakable: Apple no longer intends to treat a single in-house model stack as the center of its mobile AI story.",
					"Instead, the company is formalizing a multi-provider architecture that lets Siri and related system features route work across Apple models and selected third-party frontier systems. That is a major break from the company's historic instinct to own every important layer itself.",
				},
			},
			{
				Heading: "Apple's Integrated Model and Its Limit",
				Paragraphs: []string{
					"For decades Apple won by controlling the silicon, the operating system, the store, and the user experience as a single integrated stack. That formula depended on Apple being able to build or buy the best version of each critical layer.",
					"Frontier AI changed that assumption. Apple's on-device systems remained strong for focused tasks, but open-ended reasoning, generation, and agentic behavior advanced fastest in external labs. Apple Intelligence acknowledged the gap; iOS 27 turns that workaround into core platform design.",
				},
			},
			{
				Heading: "iOS 27 Takes the Routing Logic Further",
				Paragraphs: []string{
					"The beta centers on a new Intelligent Services Layer, or ISL, which treats on-device Apple models, cloud-extended Apple systems, and third-party providers as candidates inside one routing framework. The system chooses the model based on task type, privacy requirements, available context, and expected performance.",
					"At launch, the reported external providers are OpenAI and Anthropic. The important point is not just who is included first, but that Apple appears to have designed the interface for expansion, making model orchestration rather than model exclusivity the core product decision.",
				},
			},
			{
				Heading: "What Google Is Doing Differently",
				Paragraphs: []string{
					"That stands in sharp contrast to Android's Gemini-led approach. Google has spent the last two years binding mobile AI tightly to its own model family, from local inference up through larger cloud tasks, which gives it stronger end-to-end tuning across the OS and model boundary.",
					"Apple is making the opposite bet. A multi-provider layer may be less coherent than a single-stack design, but it is more adaptable if leadership in the model race keeps changing. In practice, Apple is positioning the iPhone as the best distribution and orchestration surface for whichever models remain strongest over time.",
				},
			},
			{
				Heading: "The Privacy Architecture That Makes It Viable",
				Paragraphs: []string{
					"Apple's challenge is not only technical but political. A system that routes user requests to outside providers only works if Apple can convincingly separate personal tasks from general ones and keep sensitive workloads inside its own trusted infrastructure.",
					"The iOS 27 design described by the beta does that through tiered classification and user controls. Personal requests stay on-device or inside Apple's private cloud path, while broader tasks such as general knowledge, creative work, or code generation can be routed outward under explicit settings.",
				},
			},
			{
				Heading: "What This Means for the Platform Wars",
				Paragraphs: []string{
					"The larger significance is architectural. Google is betting that controlling the model stack will produce the strongest mobile AI experience. Apple is betting that trusted orchestration across multiple providers will be more durable than any single lab's temporary lead.",
					"For AI companies, a place inside Apple's routing layer could become one of the most valuable distribution positions in consumer technology. For users, the result is that the mobile platform war is no longer just about devices or app ecosystems. It is now directly about which operating system manages AI best.",
				},
			},
		},
	},
	{
		Title:     "The Sprint Is Real: Inside xAI's Grok 4 Race to the Top",
		Slug:      "sprint-is-real-inside-xai-grok-4-race-to-the-top",
		Date:      "May 10, 2026",
		Tag:       "Models",
		Summary:   "xAI's Colossus cluster, X data advantage, and sub-version release tempo are turning Grok 4 into a case study in how the frontier model race has compressed into continuous iteration.",
		SourceDoc: "https://docs.google.com/document/d/1Cg_DKrFkk4eDHtiV0aMIYdOOnr7p-3csicbxgYDo2ow/edit",
		Sections: []Section{
			{
				Paragraphs: []string{
					"When Elon Musk founded xAI in the summer of 2023, skeptics had a straightforward critique: he was late to the frontier AI race, lacked a proprietary data moat, and was already stretched across several companies. Grok 1 did little to quiet that skepticism.",
					"Eighteen months later, that picture looks very different. Grok 4 made xAI a genuine top-tier contender, and the expected arrival of Grok 4.4 suggests the company is operating on a model improvement loop that is unusually fast even by current industry standards.",
				},
			},
			{
				Heading: "The Colossus Advantage",
				Paragraphs: []string{
					"xAI's acceleration starts with compute. Its Colossus cluster in Memphis began with 100,000 NVIDIA H100 GPUs and has expanded since, giving the company one of the largest dedicated AI training systems in the world.",
					"That matters because frontier progress depends not just on a single long training run, but on the ability to test architectures, tune hyperparameters, and pursue multiple experiments in parallel. A cluster built primarily for training gives xAI more freedom to iterate than competitors balancing internal compute across many product workloads.",
				},
			},
			{
				Heading: "What Grok 4 Got Right",
				Paragraphs: []string{
					"On release, Grok 4 posted competitive results against GPT-5.5 and Claude Sonnet 4.6 on reasoning tasks, while standing out more clearly on coding and long-context retrieval. Users also described it as more direct and less heavily hedged than rival systems.",
					"That style has made Grok 4 more distinctive, but also more controversial. Safety researchers continue to question its refusal behavior in sensitive domains, while xAI argues that over-cautious systems can themselves withhold useful information. Separate from that debate, Grok's integration into X gives it a real-time social data stream that no direct rival currently matches at comparable scale.",
				},
			},
			{
				Heading: "The 4.4 Sprint and What It Signals",
				Paragraphs: []string{
					"The bigger story is cadence. Instead of treating each release as a rare marquee event, xAI has moved through Grok 4.1, 4.2, 4.3, and now an anticipated 4.4, with each update targeting a narrower capability cluster such as long-context reasoning, multimodal analysis, or tool use.",
					"That shift turns frontier model development into something closer to continuous delivery. Reports around Grok 4.4 point to stronger multi-step reasoning and math performance, suggesting xAI is trying to close the remaining gaps by shipping improvements as soon as they are production-ready rather than waiting for the next large brand reset.",
				},
			},
			{
				Heading: "The Race Nobody Can Afford to Lose",
				Paragraphs: []string{
					"xAI's pace reflects a broader industry reality: any capability lead now decays quickly. GPT-5.5 Instant, Claude Sonnet 4.6, Gemini 3.1 Ultra, and the Grok 4 line are all arriving on timelines that make durable technical dominance difficult to hold.",
					"For users, that compression mostly means faster improvement. For the industry, it raises a harder operational question about whether safety review can keep pace with sub-monthly releases, especially when those systems are embedded into large public platforms and exposed to hundreds of millions of people in live contexts.",
				},
			},
		},
	},
	{
		Title:     "The CAISI Reversal: What Washington's Sudden Policy Pivot Means for AI Development",
		Slug:      "caisi-reversal-what-washingtons-sudden-policy-pivot-means-for-ai-development",
		Date:      "May 10, 2026",
		Tag:       "Policy",
		Summary:   "The White House's withdrawal of the CAISI framework removes the closest thing the US had to a moderate federal AI safety regime and leaves developers facing a widening gap between domestic permissiveness and European enforcement.",
		SourceDoc: "https://docs.google.com/document/d/10fv_JeJowwPbo9fhUBit3s2_qFcR17QM9ExCjmG56oY/edit",
		Sections: []Section{
			{
				Paragraphs: []string{
					"The White House's decision to withdraw support for the Comprehensive AI Safety and Innovation Standards framework, or CAISI, landed as a sharp reversal only weeks after it had seemed like the most viable bipartisan federal AI governance effort in the United States.",
					"The administration framed the move as a competitiveness decision, arguing that mandatory pre-deployment evaluations would slow American labs. The immediate result, though, is a policy vacuum at the federal level just as frontier systems are moving into broader commercial and agentic use.",
				},
			},
			{
				Heading: "What CAISI Was",
				Paragraphs: []string{
					"CAISI was built over two years with participation from NIST, DARPA, the FTC, and a wider civil society coalition. It was meant to create a national baseline above the growing patchwork of state-level AI rules.",
					"Its main pillars were mandatory capability evaluations for frontier models above a compute threshold, incident reporting within 72 hours for serious failures or misuse, and a voluntary certification label that companies could use to signal compliance to customers and procurement teams. By design, it was milder than the EU AI Act and left significant room for open-source development.",
				},
			},
			{
				Heading: "What Happened",
				Paragraphs: []string{
					"The immediate trigger was an April OSTP report claiming CAISI evaluations could delay qualifying model releases by four to twelve months. Critics quickly noted that the estimate relied heavily on industry-provided numbers and clashed with shorter internal modeling attributed to NIST.",
					"The political case against CAISI was straightforward: major labs argued that Chinese competitors do not face equivalent friction, so imposing mandatory pre-deployment checks in the US would create a structural disadvantage. That argument ultimately carried more weight in Washington than the case for moderate early oversight.",
				},
			},
			{
				Heading: "What Safety Researchers Are Saying",
				Paragraphs: []string{
					"Researchers who helped shape the framework argue that CAISI was intentionally designed to be lightweight enough to preserve development speed while still giving regulators visibility into serious failures. Their concern is that voluntary disclosure leaves the public and the government dependent on company judgment about what to reveal and when.",
					"That concern is amplified by recent incidents in agentic and coding systems that reportedly went unreported to federal agencies for extended periods. Under CAISI, those events would likely have triggered mandatory disclosure on a fixed timeline rather than informal, selective reporting.",
				},
			},
			{
				Heading: "The Regulatory Vacuum and the EU Divergence",
				Paragraphs: []string{
					"The reversal also widens the operational gap between the US and Europe. The EU AI Act is already enforceable for high-risk uses, which means American companies selling into Europe still have to meet binding obligations abroad even as comparable federal rules disappear at home.",
					"That asymmetry may prove more expensive than its advocates expect. Avoiding a domestic framework does not remove compliance work for global companies; it just forces them to operate across more divergent legal regimes while hoping no major incident triggers a far harsher response later.",
				},
			},
			{
				Heading: "What Comes Next",
				Paragraphs: []string{
					"The White House did not eliminate AI governance entirely. It preserved the voluntary certification track and called for a new NIST working group to pursue so-called innovation-compatible standards by late 2027.",
					"That timeline is long relative to frontier model progress. If the current pace holds, the US will reach several more generations of more capable agentic systems before a new federal baseline is even proposed, increasing the odds that future regulation arrives only after a public failure makes a slower, more deliberate path impossible.",
				},
			},
		},
	},
	{
		Title:     "Three AIs, Three Laws: Why the US, EU, and China Can't Agree on What to Do About AI",
		Slug:      "three-ais-three-laws-us-eu-china-ai-governance",
		Date:      "May 8, 2026",
		Tag:       "Policy",
		Summary:   "The EU is enforcing compliance-first rules, the US is pushing federal preemption in the name of innovation, and China is binding AI policy directly to state control, leaving global builders with three incompatible operating environments.",
		SourceDoc: "https://docs.google.com/document/d/1P5XdnlyFy6EiGAl-Q6899hnvQj1kKAV3nhf9sABQbzg/edit",
		Sections: []Section{
			{
				Heading: "The Compliance Framework in Europe",
				Paragraphs: []string{
					"The EU AI Act entered into force in August 2024 and continues phased implementation through 2027, but by 2026 the most important working pieces are already active: prohibited practices, GPAI obligations, transparency duties, and meaningful penalties.",
					"That architecture reflects a simple European premise: if AI systems will influence work, credit, healthcare, public services, and speech, then documentation, auditability, and human oversight need to be built in before scale arrives, not after.",
				},
			},
			{
				Heading: "The United States Wants Federal Uniformity",
				Paragraphs: []string{
					"The American picture is moving in the opposite direction. Instead of accepting a patchwork of aggressive state rules, Washington is leaning into preemption, arguing that the country needs one innovation corridor rather than fifty competing AI regimes.",
					"The logic is straightforward even if the politics are not: if frontier AI is a strategic industry, then the federal government does not want California, New York, and Texas each imposing a different rulebook on model development. Innovation speed is being treated as a national asset.",
				},
			},
			{
				Heading: "China Treats AI as State-Integrated Infrastructure",
				Paragraphs: []string{
					"China's model is not just tighter regulation in the abstract. It embeds AI governance into the state's existing content-control and administrative systems, with rules on labeling, complaint handling, consent, and synthetic media all operating inside a larger project of political legibility.",
					"For developers, that means compliance is less about proving abstract fairness and more about proving controllability. Companies can move quickly, but only within guardrails that remain clearly aligned to state priorities.",
				},
			},
			{
				Heading: "The Regulatory Trilemma",
				Paragraphs: []string{
					"Together these frameworks create a real deployment problem. No single architecture fully satisfies all three blocs at once, which means multinational AI products increasingly need region-specific data flows, compliance documents, and even model behaviors.",
					"The important technical takeaway is that regulation is now part of system design. Teams that once treated policy as an afterthought are going to find that where they deploy AI increasingly shapes what they are allowed to build at all.",
				},
			},
		},
	},
	{
		Title:     "From Text-to-Video to Intent-to-Video: The Quiet Revolution in AI Filmmaking",
		Slug:      "from-text-to-video-to-intent-to-video-ai-filmmaking",
		Date:      "May 8, 2026",
		Tag:       "Media",
		Summary:   "The newest video models are moving beyond clip generation toward systems that understand pacing, continuity, sound, and narrative purpose, turning prompt boxes into early-stage directing tools.",
		SourceDoc: "https://docs.google.com/document/d/1nXzDuufjjF9om2ATFOtnTzRpoRkgu8J0fzhqKv4ceEM/edit",
		Sections: []Section{
			{
				Heading: "Seedance 2.0 and the Multimodal Turn",
				Paragraphs: []string{
					"ByteDance's Seedance 2.0 is a good example of where the market is moving. It accepts text, images, audio, and video together, and it gives users tighter control over continuity, movement, and scene coherence instead of forcing them to regenerate clips until something usable appears.",
					"That matters because creators do not think in isolated prompts. They think in sequences, references, camera moves, and emotional timing. The winning tools are the ones that can absorb that production context rather than simply rendering a sentence.",
				},
			},
			{
				Heading: "Why Fidelity Is No Longer the Only Metric",
				Paragraphs: []string{
					"Runway's latest systems continue to lead on visual fidelity and temporal consistency, but the deeper shift is that sharp images are no longer enough. Production teams need characters that stay consistent across cuts, environments that do not drift, and revisions that preserve what already worked.",
					"A beautiful shot is still marketing. A coherent scene is what makes the tool usable. That is why continuity is becoming the competitive battlefield in AI filmmaking.",
				},
			},
			{
				Heading: "Google's Omni Bet: Sound and Vision Together",
				Paragraphs: []string{
					"Google's coming Omni model points to another major change: native synchronization between sound and image. Instead of adding audio after the fact, the system is expected to reason about the relationship between thunder, motion, distance, and ambience as part of the same generation task.",
					"That makes AI video feel less like image synthesis with motion and more like scene synthesis. It is a fundamentally different product category because it begins to model experience, not just frames.",
				},
			},
			{
				Heading: "Intent-to-Video Is an Architectural Shift",
				Paragraphs: []string{
					"The phrase 'intent-to-video' captures what is changing under the hood. Older models were trained to respond to descriptions. Newer systems are being shaped to infer the sensory and emotional outcome a creator wants, then translate that into framing, pacing, lighting, and motion.",
					"That is why this revolution feels quiet. The public still sees text boxes, but the real progress is in timeline awareness, cross-shot memory, and systems that understand what creators mean when they ask for tension, restraint, or a delayed reveal.",
				},
			},
		},
	},
	{
		Title:     "Google's 75% Stat is the Wake-Up Call Software Engineers Needed",
		Slug:      "googles-75-percent-stat-wake-up-call-software-engineers-needed",
		Date:      "May 8, 2026",
		Tag:       "Engineering",
		Summary:   "Google's claim that AI is generating more than 75% of some new code paths is less a flex than a signal that engineering value is shifting toward design, review, testing, and operational judgment.",
		SourceDoc: "https://docs.google.com/document/d/18CSkDTDZreSOkPc45uBKHlq4M34tXai54-kWxlVplDE/edit",
		Sections: []Section{
			{
				Heading: "The Number Matters Because of Who Said It",
				Paragraphs: []string{
					"When a company the size of Google says AI is involved in more than 75% of some new code output, the exact percentage is less important than the direction of travel. Mature software organizations are actively reorganizing around AI-assisted production.",
					"The wake-up call is not that code is disappearing. It is that raw code generation is becoming cheaper, which moves more engineering value into architecture, system boundaries, and the ability to decide what should actually ship.",
				},
			},
			{
				Heading: "MCP and the Agentic Stack Are Becoming Foundational",
				Paragraphs: []string{
					"The bigger story is that autonomous coding is becoming infrastructure. The Model Context Protocol has turned into a common control plane for tools and agents, while large-context models are increasingly designed to write, run, test, and revise code inside the same loop.",
					"That changes the day-to-day job. Engineers are no longer just using autocomplete. They are starting to supervise small software systems that can act on repositories, test suites, shells, and documentation with increasing independence.",
				},
			},
			{
				Heading: "Output Is Cheap, Judgment Is Not",
				Paragraphs: []string{
					"AI can already produce scaffolding, refactors, tests, and glue code at a pace that changes team throughput. What it still struggles to replace is judgment: understanding failure modes, making tradeoffs under ambiguity, protecting security boundaries, and keeping systems maintainable over time.",
					"That means the defensible engineering skill is not just prompting. It is knowing how to turn fast machine output into trustworthy software under real production constraints.",
				},
			},
			{
				Heading: "What Engineers Should Do Now",
				Paragraphs: []string{
					"The practical response is to use agentic tooling in real repositories and learn where it breaks. Teams need people who can define review standards, shape evaluation loops, and catch subtle regressions that high-speed generation tends to hide.",
					"If AI keeps raising raw code throughput, the engineers who become more valuable will be the ones who can reason about systems, product intent, and operational risk while directing a growing fleet of coding agents effectively.",
				},
			},
		},
	},
	{
		Title:     "A Broken Promise Worth $134 Billion: The OpenAI Trial Putting AI Governance Under Oath",
		Slug:      "broken-promise-worth-134-billion-openai-trial-ai-governance-under-oath",
		Date:      "May 6, 2026",
		Tag:       "Governance",
		Summary:   "The Musk v. Altman case is exposing the most detailed public record yet of how OpenAI's nonprofit mission, commercial structure, and power politics collided under oath.",
		SourceDoc: "https://docs.google.com/document/d/1h4FGSY_qMo8ISTh44AmaH5upQs4pzo67WUb-GBdLrh0/edit",
		Sections: []Section{
			{
				Paragraphs: []string{
					"The most important thing about the Musk v. Altman trial is not who wins. Courts are skeptical institutions and $134 billion judgments are rare. What matters is what the trial is producing: the most detailed public record of how the world's most influential AI lab was actually built, governed, and eventually transformed from a nonprofit into a $157 billion commercial entity.",
					"That record is being made under oath, in a federal court in Oakland, California, and it is not flattering to anyone involved.",
				},
			},
			{
				Heading: "The Case in Brief",
				Paragraphs: []string{
					"Elon Musk sued OpenAI, Sam Altman, and Greg Brockman in 2024, claiming that OpenAI broke its founding promise to remain a nonprofit developing AI for the benefit of humanity. Musk contributed approximately $38 million to OpenAI in its early years and argues those funds were solicited under specific commitments about the organization's mission and structure.",
					"The relief sought is extraordinary: up to $134 billion in damages, the removal of Altman and Brockman from their positions, and the unwinding of OpenAI's conversion to a Public Benefit Corporation, completed on October 28, 2025 with regulatory approval from the attorneys general of California and Delaware.",
					"Under the new structure, the OpenAI Foundation holds 26%, Microsoft holds 27%, and employees and investors hold the remaining 47%. Sam Altman reportedly took no equity in the restructuring.",
				},
			},
			{
				Heading: "The Bombshells From the Stand",
				Paragraphs: []string{
					"Musk testified for three days in week one. He said he was duped. He warned the court that AI could kill humanity. And then he admitted under oath that xAI, his own AI company, distills models from OpenAI's outputs, training Grok on the very outputs of the system he is suing to stop.",
					"Musk also texted Greg Brockman to explore settlement two days before the trial began. OpenAI rejected it. That overture suggests Musk's legal team had genuine doubts about their position, which aligned with prediction market pricing on May 5 that implied roughly a 22% chance of a Musk win.",
					"In week two, Brockman flatly disputed Musk's version of OpenAI's founding. He testified that he had made no commitments to Musk about corporate structure and heard no one else do so. He also said Musk secretly redirected several OpenAI employees in 2017 to spend months overhauling Tesla's Autopilot program without OpenAI's knowledge or compensation.",
				},
			},
			{
				Heading: "What This Trial Is Actually Deciding",
				Paragraphs: []string{
					"Beyond Musk's personal claims, the trial is testing whether the founding mission of a nonprofit AI organization creates legally enforceable constraints on how that organization commercializes its work.",
					"If a court rules that OpenAI made binding legal commitments that the Public Benefit Corporation restructuring violated, the implications extend far beyond this case. Dozens of AI and technology organizations have hybrid nonprofit-commercial structures, and every one of them would face similar scrutiny.",
					"The Microsoft dimension adds another layer. Microsoft holds 27% of OpenAI's commercial entity and has tied Copilot, Azure AI, and its broader product stack to that relationship. Any ruling that forced a restructuring would reverberate through one of the world's most valuable companies, not just OpenAI.",
					"Shivon Zilis was identified in the draft as the next witness. Because she served on OpenAI's nonprofit board from 2019 to 2023, her testimony could either corroborate or contradict Musk's account with board-level knowledge of the conversion path.",
					"However the ruling lands, the trial has already succeeded in elevating the question of who should control AI from philosophy to legal precedent. The founding tensions between safety mission and commercial reality are now documented in a public court record that the industry will keep citing long after the damages question is resolved.",
				},
			},
		},
	},
}

// Posts returns all published posts in publication order.
func Posts() []Post {
	return clonePosts(posts)
}

// FindBySlug returns a published post when the slug matches exactly.
func FindBySlug(slug string) (Post, bool) {
	for _, post := range posts {
		if post.Slug == slug {
			return clonePost(post), true
		}
	}

	return Post{}, false
}

func clonePosts(src []Post) []Post {
	dst := make([]Post, len(src))
	for i, post := range src {
		dst[i] = clonePost(post)
	}

	return dst
}

func clonePost(src Post) Post {
	dst := src
	dst.Sections = make([]Section, len(src.Sections))
	for i, section := range src.Sections {
		dst.Sections[i] = Section{
			Heading:    section.Heading,
			Paragraphs: append([]string(nil), section.Paragraphs...),
		}
	}

	return dst
}
