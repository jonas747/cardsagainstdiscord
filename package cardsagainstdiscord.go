package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "Brookly Nine Nine",
		Description: "Brookly Nine Nine TV Show Themed",
		Prompts: []*PromptCard{		
            &PromptCard{Prompt: `That's how we do %s in the Nine-Nine.`},
			&PromptCard{Prompt: `Peralta, what are you doing? %s `},
			&PromptCard{Prompt: `I once got arrested for %s `},
			&PromptCard{Prompt: `The worst thing to say to a cop is %s`},
			&PromptCard{Prompt: `The worst thing do in front of a cop is %s`},
			&PromptCard{Prompt: `The best undercover name is %s`},
			&PromptCard{Prompt: `The most disturbing thing I've seen as a cop is%s`},
			&PromptCard{Prompt: `The most bad ass catchprase is %s`},
			&PromptCard{Prompt: `Someone painted %s on our minivan`},
			&PromptCard{Prompt: `Terry loves %s`},
			&PromptCard{Prompt: `Eyes up here, Gina. I'm more than %s`},
			&PromptCard{Prompt: `19 milk less years I have gone, and for some reason I can't quit %s?`},
			&PromptCard{Prompt: `%s, from %s. Sounds like a fun treat.`},
			&PromptCard{Prompt: `Oh, I'm as serious as %s`},
			&PromptCard{Prompt: `Amy Santiago was actually late to work because of %s`},
			&PromptCard{Prompt: `When will you quit police work and pursue your dream of %s?`},
			&PromptCard{Prompt: `I'm making a scrapbook as a present for Capt. Holt. It's titled %s`},
			&PromptCard{Prompt: `They do call me Tenacious Ray down at the country club. Because for the past 10 years, I've been suing them for %s`},
			&PromptCard{Prompt: `The most intimate thing you can do to your lover is %s, other than washing their hair`},
			&PromptCard{Prompt: `No, from now on call me, %s.`},
			&PromptCard{Prompt: `Ew, it's like dating %s.`},
			&PromptCard{Prompt: `I know, it's so hot. %s`},
			&PromptCard{Prompt: `These silences are very hard for me to read. What's the vibe in here? %s`},
			&PromptCard{Prompt: `We're %s, Charles!`},
			&PromptCard{Prompt: `The alcohol has rendered me %s.`},
			&PromptCard{Prompt: `What is it about me that screams loser? %s`},
			&PromptCard{Prompt: `There's nothing more intoxicating than %s`},
			&PromptCard{Prompt: `I've been asked to deliver a message of hope. This is what I have so far: %s.`},
			&PromptCard{Prompt: `The reason for Jake's financial issues is %s`},
			&PromptCard{Prompt: `Is that Mr. MacLaine I hear? Cause someone's about to %s`},
			&PromptCard{Prompt: `Seriously, what happened to Hitchcock and Scully? %s`},
			&PromptCard{Prompt: `The hardest case I've ever solved involved %s`},
			&PromptCard{Prompt: `My plan to win the Halloween Heist is simple. It will involve %s`},
			&PromptCard{Prompt: `Who is the real Amazing Detective/Human/Genius? %s`},
			&PromptCard{Prompt: `The only thing worse than failing to catch a criminal is %s`},
			&PromptCard{Prompt: `The night shift is keeping you and Kevin apart. You two just need to %s.`},
			&PromptCard{Prompt: `Captain banned %s from the office due to the Gina incident.`},
			&PromptCard{Prompt: `She sabotaged my career because I refused to %s.`},
			&PromptCard{Prompt: `Dig deep, think of something scary %s`},
			&PromptCard{Prompt: `I referred to myself as %s to several beautiful Latina ladies.`},
			&PromptCard{Prompt: `The Nine-Nine hates the Fire Department because of %s`},
			&PromptCard{Prompt: `Here are your drinks, and %s from Mr. Boyle`},
			&PromptCard{Prompt: `Good point. Rosa would want a much bigger %s.`},
			&PromptCard{Prompt: `Here at the Fun Zone, we live by one rule. When it's your birthday you're always %s.`},
			&PromptCard{Prompt: `Gina Linetti has a new dance troupe called %s.`},
			&PromptCard{Prompt: `%s are dumb.`},
			&PromptCard{Prompt: `My kids think their pre-school teacher's %s.`},
			&PromptCard{Prompt: `Tiny Terry loves his %s.`},
			&PromptCard{Prompt: `I need you to unleash %s`},
			&PromptCard{Prompt: `I can't hear you Kevin, you're %s.`},
			&PromptCard{Prompt: `Capt. Holt should name his car %s instead of Gertie.`},
			&PromptCard{Prompt: `Pimento has %s disease.`},
			&PromptCard{Prompt: `- I don't want to say what using good vocabulary gets me from Amy: %s.`},
			&PromptCard{Prompt: `Get familiar with %s or get familiar with the unemployment line.`},
			&PromptCard{Prompt: `No, I don't wanna listen Boyle. You are trying to %s and that will not fly!`},
			&PromptCard{Prompt: `What does the Ebony Falcon do? %s`},
			&PromptCard{Prompt: `You've lost the ability to %s. You just plain boring.`},
			&PromptCard{Prompt: `The night that you %s for 20 seconds and I became obsessed with you forever`},
			&PromptCard{Prompt: `I'll %s to perk up my little man`},
			&PromptCard{Prompt: `My dreams are coming true. You and me getting %s together.`},
			&PromptCard{Prompt: `Up is up. Down is down. The Vulture still %s.`},
			&PromptCard{Prompt: `What's your happy place? %s`},
			&PromptCard{Prompt: `Say %s again and I will shoot you in the stomach.`},
			&PromptCard{Prompt: `He says %s just like you.`},
			&PromptCard{Prompt: `You call that %s?`},
			&PromptCard{Prompt: `Fine, but we didn't have to endure this crap when we danced in %s.`},
			&PromptCard{Prompt: `Your dress makes you look like %s.`},
			&PromptCard{Prompt: `Damn it, %s! You sold me out.`},
			&PromptCard{Prompt: `%s got served.`},
			&PromptCard{Prompt: `You embarrassed yourself in front of %s.`},
			&PromptCard{Prompt: `The ghost of the %s haunts the precinct!`},
			&PromptCard{Prompt: `She had a replacement hip with some serious %s`},
			&PromptCard{Prompt: `It's like %s with a transformer`},
			&PromptCard{Prompt: `Kind, sober and %s`},
			&PromptCard{Prompt: `Time for me to get out there and %s`},
			&PromptCard{Prompt: `It's not your fault, I was %s`},
			&PromptCard{Prompt: `It was slightly less %s with you.`},
			&PromptCard{Prompt: `Everytime you shake me, I see it %s.`},
			&PromptCard{Prompt: `My mother has a fantastic %s.`},
			&PromptCard{Prompt: `I %s this stewardess in Sweden once.`},
			&PromptCard{Prompt: `And sir, nothing would make me %s than being your big daddy.`},
			&PromptCard{Prompt: `Something no lover of yours has ever %s.`},
			&PromptCard{Prompt: `I got a nice soft %s so I don't intimidate the husbands.`},
			&PromptCard{Prompt: `Oh honey, that flat ass is perched right on top of my %s.`},
			&PromptCard{Prompt: `Because I was your %s and your lover? Sorry I just can't control my %s around you.`},
			&PromptCard{Prompt: `The only way out of this hole is to %s.`},
			&PromptCard{Prompt: `I can picture you as a teenager just %s all over New York City.`},
			&PromptCard{Prompt: `No hole's too tight for these %s tips.`},
			&PromptCard{Prompt: `There's %s in our genes.`},
			&PromptCard{Prompt: `I just got it out of the vent to rub it in your %s`},
			&PromptCard{Prompt: `Don't be shy, just say what you will do international sex symbol %s's body.`},
			&PromptCard{Prompt: `Bill is %s. He's able to use more than just his mouth.`},
			&PromptCard{Prompt: `There's not even any soft parts in the middle we can %s`},
			&PromptCard{Prompt: `Everybody stop %s! We just had harassment training.`},
			&PromptCard{Prompt: `There are no %s in Scottdale, Margo.`},
			&PromptCard{Prompt: `Changed my relationship status to %s. Pony up y'all!`},
			&PromptCard{Prompt: `I came down with a big old %s infection.`},
			&PromptCard{Prompt: `Ain't nothing but %s.`},
			&PromptCard{Prompt: `Teddy is %s people out of the competition.`},
			&PromptCard{Prompt: `You've been in there like five times. What are you gonna do, annoy him into %s?`},
			&PromptCard{Prompt: `Do you need a %s tutor because the department will provide you one.`},
			&PromptCard{Prompt: `Dad said you were giving us tickets to the %s movie.`},
			&PromptCard{Prompt: `Sorry sir that no one here wants to %s, you dusty old %s.`},
			&PromptCard{Prompt: `This is a direct order. Stop %s.`},
			&PromptCard{Prompt: `Looks like a sticky %s.`},
			&PromptCard{Prompt: `When I was alone in her office, I changed her autocorrect. When she types in Wuntch, it gets replaced with %s.`},
			&PromptCard{Prompt: `Where's Jake? Did you %s him, Terry?`},
			&PromptCard{Prompt: `Thank you for %s my feelings.`},
			&PromptCard{Prompt: `The magic of %s. I'm			&PromptCard{Prompt: ` in a state of total euphoria.`},
			&PromptCard{Prompt: `Hey-hey! New %s alert! `},
		},
		Responses: []ResponseCard{
			`Making a badass entrance and realizing you have fudge on cheek`,
			`Your subordinate tells you to bone your husband`,
			`A cockroach disguised as an old leather chair`,
			`The satisfaction of finally defeating Wario`,
			`Use Slut Sauce as a bulletproof vest`,
			`Masturbating to a perfect school attendance record`,
			`Fart uncontrollably while lifting a car`,
			`Winning the Halloween Heist`,
			`Getting the truth out of someone by playing the guitar and screaming`,
			`Play the gay card and say YAS QUEEN`,
			`A dry boy is a smart boy`,
			`Have a feud with a police horse`,
			`The clear absence of a penis`,
			`A pair of thick weighty breasts of a female woman`,
			`Giving the B a C in her A`,
			`An engaged man distributing STDs to his coworkers`,
			`A police captain talks about erotic barrels`,
			`Calling police weapons adult toys`,
			`Abu Dhabi? Abu don't bother`,
			`Gina Linetti Spaghetti Confetti`,
			`Scully's foot fungus cream`,
			`Mentioning your anal canyon during a police briefing`,
			`Operation Beans`,
			`Reno Vegas and Tex Dallas`,
			`I NEED MY MUSHU PORK`,
			`Responsible agricultural practices`,
			`Ms Pac-Man's nipple`,
			`Scrotal Recall`,
			`Papa! There's a monster in my closet`,
			`Welcome to the party pal`,
			`YIPPIE KAYAK OTHER BUCKETS`,
			`Tarantulina Jolie`,
			`Cause a bomb panic in order to give your boss a Christmas present`,
			`Feeling trepidation at the prospect of a parentless existence`,
			`Speed up your gay wedding because you are unsure if it will stay legal`,
			`I'm off to plow my mistress`,
			`COWABUNGA MOTHER`,
			`KUDOS TO OUR BRAVE OFFICERS`,
			`Spelling case with a K`,
			`Meeting your wife at an orgy`,
			`An award for "Most likely to befriend a school administrator"`,
			`Say "those slacks are a knockout" to your boss at his birthday party`,
			`The human form of 💯`,
			`Experimenting with an unfamiliar acronym in public`,
			`A white woman in prison in Texas`,
			`My lint is oblong. My lint is blue.`,
			`Sniffing a Conflict Resolution binder for pleasure`,
			`Funky Cats and their Fiesty Stats`,
			`Teaching father the math`,
			`Thoughts on Beyonce's lemonade`,
			`Did someone order a tracheotomy?`,
			`BLACK PEOPLE CAN SELL DRUGS`,
			`Wearing a diaper to a test`,
			`My fantasy threesome`,
			`Taking some weight off your cloven hooves`,
			`You have a little bit of goop in your eye`,
			`A straight, white police commissioner`,
			`Wearing a pantsuit as a disguise`,
			`Eating sticks and stones for breakfast`,
			`A cop negotiating in order to get back his best friend's sperm`,
			`Gummy bears wrapped in a fruit roll up`,
			`Wet burrito`,
			`The touch of a woman in many moons`,
			`A Snaccident`,
			`The Fellowship of the Key`,
			`We guard what you lick`,
			`TV and Cake were my parents`,
			`Mayo nut-spoonsies`,
			`Make your boobies do the bouncy thing`,
			`Adam Sandler runs an auction to stuff a sock inside a cop's mouth`,
			`Efficiency. Efficiency. Efficiency.`,
			`Biting a guy's butt off at a WNBA game`,
			`Hilary Clinton and Kim Jong-Un make out in a holding cell`,
			`Getting makeup tips from a prostitute in the holding cell`,
			`Undiscovered muscle`,
			`Amount of fudge in a calorie`,
			`Making desk yogurt at work`,
			`Daniel Craig hands + close up`,
			`Calling Thanksgiving Turkey Day`,
			`A cop hitting Santa on the testicles in front of children`,
			`Playing Wife or Dog at work`,
			`The Fire Department`,
			`Singing Pina Coladas in front of a dead body`,
			`A police captain hula-hooping and spraining his wrist`,
			`Dry meatballs`,
			`Pepper-aloni pizza`,
			`Yell HOT DAMN after winning an office bet`,
			`Sending a condolences email to your captain as 'My Stinky Butt'`,
			`Wunch Time is Over!`,
			`Three people with pants on having a normal conversation`,
			`Pee on the whole city from the top of the Empire State Building`,
			`Rubber band ball`,
			`May I have some cocaine?`,
			`A perp looking all perpy perping his way down perp street`,
			`MR. GRAAAAAAPESSSS`,
			`Smooshing Booties`,
			`Misdemeanors weiners`,
			`Tommy Gobbler`,
			`Snacky Chan`,
			`Food stain buds`,
			`Inheriting 1 million shares of Blockbuster`,
			`Two cops getting engaged while chasing a criminal`,
			`Operation 225641441636324`,
			`Discussion of rice in front of a guy in a coma`,
			`Chocolate is the devil's carob`,
			`Carob is Satan's raisin`,
			`Delicious untoasted bread`,
			`Corn Dog pewned by Go-Kart`,
			`Marco, VAMONOS!`,
			`I do a lot of meth`,
			`Kissing Vanilla Ice`,
			`This has-been has been with yo mama all week`,
			`Dianne Weist Infection`,
			`Strike Team Thunder Kill Alpha: Hard Target`,
			`S'd in the B`,
			`Creepily singing songs from the great American Songbook`,
			`Smashing the microwave using a police baton`,
			`Ignatius Pennyfeather the 9th`,
			`I AM WAGGING BRENDA`,
			`To sex with Amy!`,
			`A police captain sneaking into their detectives' bedroom and eating all of their eggs`,
			`Doing 'The Worm' while your co-workers are mourning the death of your old boss`,
			`Using the main hole or no hole and choosing no hole`,
			`The First Pubic War, the sexiest of all wars`,
			`The Greeks climbing out of the Trojan horse's butt`,
			`Praying for a hostage situation in the hopes of becoming a negotiator`,
			`Making suspects sing 'I Want It That Way' and forgetting one of them is a murderer`,
			`Cops abusing their power in order to get a veil dry-cleaned quickly`,
			`Rant in front of your co-workers with your penis hanging out`,
			`Calling your partner's butt 'da bomb' in your wedding vows`,
			`Labia`,
			`Chanting "Broken Penis" during a police briefing`,
			`A police captain orders his best detective to prevent the worst detective from spilling condiments on himself`,
			`Nature's bullseye`,
			`Dan Dan, The Enema Man!`,
			`A captain punishing his employee for tardiness by high-fiving every other employee`,
			`Mama Magglione`,
			`The G-Hive`,
			`A mariachi band in a police precinct`,
			`Stress-eating egg yolks`,
			`Call your husband during a morning briefing to tell them how you got pranked by your employees`,
			`Kidnapping a useless cop to defeat a dirty cop`,
			`Revenge on Dorothy for killing your sister`,
			`Raw chicken is healthier to eat`,
			`A cop talking about his pelvis not knowing that there has been assassination attempt`,
			`The Boyhunter`,
			`Sighing during the whole length of an elevator ride`,
			`Accusing Brian of stealing cocaine because he has a hot tub`,
			`Your wife has an affair with her hairdresser, you film it and threaten to release the tape. She says she doesn't care and put it on the internet herself. When it starts making money, you sue her for half the profits.`,
			`A Korean toilet ghost`,
			`Rapping your resume to your boss`,
			`Anal-ment`,
			`Talking about getting a 7-pound mash removed from your abdomen during a briefing`,
			`Celebrating the death of your rival by throwing bagels at your coworkers and yelling BAGEL`,
			`Listening to music on your headphones while many cops are struggling to arrest a criminal inside the precinct`,
			`Flirting with a girl for 20 seconds and becoming obsessed with her forever`,
			`Debating the gross thing on an old person's chin during a police briefing`,
			`Panick and shoot a mannequin`,
			`Dropping your muffin, hitting your head while trying to grab the muffin and simultaneously stepping on your muffin`,
			`Detective Right All the Time and Detective Terrible Detective`,
			`Painting weiners on cop cars`,
			`Tasering a cantaloupe`,
			`Banana Fartman`,
			`Yes sir, I will make better mouth`,
			`Screaming at a police lineup because you don't want to miss the farmer's market`,
			`A police sergeant with huge muscles cannot understand what kind of castle has wheels`,
			`Making fat jokes during an investigation`,
			`Pretending to be dead while having sex`,
			`Don Flan`,
			`A mild-mannered doorman gets bitten on the penis by a radioactive spider and becomes the world's greatest lover.`,
			`A Crispy Mother Werewolf or Cowboy Mustard, Oslo, Norway`,
			`FOR TOO LONG, WE'VE BEEN PUT DOWN, RIDICULED MADE TO WEAR TIES. BUT NO MORE. FOR TODAY, WE DEFEAT HIM!`,
			`The Medal of Valor`,
			`The Dagger`,
			`The Hammer`,
			`The Hall Monitor`,
			`The Deuce`,
			`Fingers has grabbed the package`,
			`A baditude.`,
			`The Death Watch`,
			`Frans Bruggen playing the flute`,
			`The cleavage cloaks the camera with its curves`,
			`Operation: Oh crap, wrong vent`,
			`Al The Janitor`,
			`Alpha Team, This is the Golden Sparrow. Rendezvous at Drop Zone Yankee.`,
			`I will slit you both open from mouth to anus and wear you like jackets`,
			`The Caboo-dale`,
			`My eyes will be glued to that ass`,
			`Put on a smile, Pork Chop`,
			`Now boy, it's time to make daddy proud`,
			`Linetti, Lin-out-ti`,
			`Hook, line and sphincter`,
			`A police captain watching a video of an American gymnast whose leotard rips and his butt gets exposed. It is highly erotic`,
			`You're just some common bitch`,
			`A police captain threatens to neuter his coworkers because he thinks someone took his dog.`,
			`Blessed be the fruit, baby.`,
			`Feeding GPS trackers in secret to your coworker in order to win a competition`,
			`Fake proposing to someone after beating them in a bet, fake proposing again in a Halloween competition and then proposing for real 2 years later in the same Halloween competition`,
			`Sergeant Dum Dum`,
			`A police captain and a sergeant insult each other's sexual activity`,
			`Convince your husband you're pregnant in the hopes of winning a competition`,
			`Tase your husband to steal a medical bracelet that belongs to a man who is at risk of brain damage`,
			`Tase your wife twice to steal a medical bracelet that belongs to a man who is at risk of brain damage`,
			`Ship your friend to New Jersey in a box to steal a medical bracelet that belongs to a man who is at risk of brain damage`,
			`Smash windows and destroy government property just for the sake of a competition`,
			`Squirt more lube and help me yank`,
			`IT'S ONE BUND TO NONE, SON`,
			`You duplicitous bitch`,
			`I'm gonna drown you in your own blood`,
			`I'm gonna rip your own arms off and beat them to death with it`,
			`I'm going to slice your Achilles' tendons, peel of your fingernails, and stick knitting needles in your eyes`,
			`Set up a surveillance system in your employees' apartment and watch them sleep just for the sake of winning a competition`,
			`Listen to your husband's therapy just for the sake of winning a competition`,
			`Hire an actor as a therapist for your husband just for the sake of winning a competition`,
			`A Thicc King`,
			`Thaboo wearing the Infinitude Gobbler in the movie Avengaboys in order to get the Infinitude Gems`,
			`PB&J: Penis, Brain and Jaw`,
			`Champagne, Mountain Range, Hugs`,
			`Imagine a letter had unprotected sex with a phone`,
			`Every Individual Gets Crayons After Telling His Aggressive Little Mongoose Painter Called Ernest Some Lies About Tiny Panda Heads. Maybe, One Kid Could Take Her Elephant Into California Except....`,
			`www.ladiesgoodhealthmag.com/sex-relationships/867599904/9432&20.html`,
			` Raymond Holt in hexadecimal code`,
			`The culinary equivalent of unprotected sex`,
			`Of course you didn't want to burst my bubble`,
			`We did this. Our sex made this happen.`,
			`Knocking Boots`,
			`Bone Bros`,
			`Workplace Bone Buds`,
			`Kwazy Cupcakes`,
			`Dumpy Chuck Norris`,
			`Dumpy Ron Weasley`,
			`Gay Robin Hood`,
			`Homeless Troll Doll`,
			`Chasing a criminal in your wedding dress`,
			`Okay muchachos, let's roll up on these muchachos`,
			`Rhonda Shawarma`,
			`Wet Blanket`,
			`Velvet Thunder`,
			`Funktown State University`,
			`RELEASE YOUR SWEETS!!`,
			`Enjoy your blood yogurt`,
			`The Chamber of Asses`,
			`The 99th Precinct`,
			`The 69th Precinct`,
			`Shoulder Nova`,
			`Vanessa Santiago`,
			`Month old Chinese food from the fridge`,
			`That's gonna leave a mark!`,
			`The Sloppy Jessica`,
			`Cantaloupe, Yes I Can!`,
			`Two married cops doing a structured debate about having kids with their boss as a moderator in a hospital`,
			`Offering Meat Pile to Sex Crimes Unit in exchange for getting 4 kilobytes of Internet`,
			`Squash's Unhinged Lunatic`,
			`You butternuts ready to get squashed?`,
			`Ray-Ray and the Beast`,
			`Auntie Ro-Ro`,
			`Uncle Chi-Chi`,
			`Better get some corticosteroids to treat that laryngeal fracture`,
			`A dirtbag is a useful part of a vaccuum.`,
			`Joke Peralta`,
			`MAN DID CRIME`,
			`Sal's Pizza`,
			`YOU'VE BEEN BOONED`,
			`New York's best at spraying stuff with water`,
			`Mouth-feel`,
			`Tell your boss he looks beautiful`,
			`MY WIFE WAS MURDERED BY A MAN IN A YELLOW SWEATER`,
			`Fire your employee because she tells you that your math is wrong`,
			`VIVA THE DOWNSTAIRS PEOPLE!`,
			`Evaluation pants`,
			`Dead guy sex`,
			`An old detective drowns in a bowl of water while taking a nap`,
			`An old detective gets strangled by his shoe laces while taking a nap`,
			`Hairy Pop-tart`,
			`Elvis Stojko, the Canadian skater figure`,
			`TWO PRANKS FOR THE PRICE OF ONE!`,
			`YA'LL JUST DRANK CEMEEEENNNTTT`,
			`Prankster prevents a police precinct from being shut down`,
			`LIKE YEAST`,
			`A police captain acts dead in an attempt to portray Elvis during a game of charades`,
			`Jeffords Junction`,
			`Mount Terry`,
			`Señor Tickle`,
			`Fill out an official police complaint in crayon`,
			`Victor Emanuel III`,
			`A talking raisin`,
			`Change the auto-correct of your rival so that her surname auto corrects to 'Butt'`,
			`Black NASA`,
			`LeBron James's new school for black astronauts`,
			` Giving each other road head`,
			`White pubes induced as a result of stress from a trial`,
			`A good old-fashioned suck-off`,
			`Fingerholes Bowling Alley`,
			`Finger Queen`,
			`Orgasm Juice`,
			`Lady Dipstick, Mr. Fib and The Milkman`,
			`A tattoo on my forehead that says 'Sinbad lives in my building'`,
			`The All-ages piano lounge`,
			`Mischief managed`,
			`Glorified EMTs who sleep in bunk beds`,
			`Clothes hanging off a genderless blob`,
			`Don Johnson it`


		},
	}

	AddPack(pack)