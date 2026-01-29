## Exercice 1

1. Quels index sont utilisés dans /tasks (GET) et pourquoi ?

Nous utilisons le même index que celui utilisé dans la base de données. Lorsque nous créons une task, l'index est attribué automatiquement et est unique. 

2. Expliquez l’importance de idx_parent dans la route API /tasks/:id/subtasks ?

Elle est utilisée comme foreign key et permet à la base de données d'établir très rapidement l'association entre les tasks et les subtasks.

3. Comment améliorer la recherche de tâches par status et due_date dans une même
requête ? et pourquoi ?

Nous souhaitons obtenir toutes les tasks ayant une date d'échéance donnée, car cela nous permet d'établir des priorités. Nous pouvons rendre cette opération plus efficace en segmentant la table selon la colonne due_date. Ainsi, lorsque nous interrogeons toutes les tasks avec la même due date donnée, elles seront toutes regroupées.

## Exercice 2

1. Qu’est-ce qu’une fonction réentrante ?

Une fonction réentrante est une fonction qui, si elle est interrompue, ne perturbe pas le reste du système et peut être exécutée à nouveau. Elles n'utilisent que des variables locales.

2. Quelle est la différence entre un thread, un fork et une coroutine ?

Thread: Processus qui s'exécute dans un cœur de processeur physique ou logique unique. Les applications multithread peuvent utiliser plusieurs cœurs de processeur pour obtenir un parallélisme et être plus rapides.

Coroutine: Processus qui s'exécute simultanément au sein du thread. Cela signifie qu'il est asynchrone, mais pas parallèle. C'est pourquoi les applications qui ne peuvent s'exécuter que dans un seul thread (par exemple, les applications NodeJS) peuvent tout de même utiliser des fonctions asynchrones. Cela est utile pour ne pas interrompre le programme lorsque nous devons attendre une réponse d'une ressource externe, comme une API ou une base de données.

Fork: Copie d'un processus existant, isolée.

3. Qu’est-ce que HELM ?

HELM est un ensemble d'outils qui vise à organiser les configurations Kubernetes. Avoir beaucoup de fichiers .yaml peut être fastidieux, HELM simplifie ces problèmes en unifiant les déploiements. HELM fournit également d'autres outils tels que la gestion des versions et les restaurations. Il dispose d'une interface CLI.

4.​ Que pensez-vous des design patterns ? (donner des exemples) Avez-vous déjà
utilisé un ou plusieurs patron(s) de conception ?

Les design patterns sont des modèles très utiles pour résoudre des problèmes courants. Voici quelques-uns de ceux que j'ai déjà utilisés :

Iterator: Très utile pour traiter de grandes quantités de données, quelle que soit la forme de leur structure.

Singleton: Utilisé pour ne pas créer de nouvelles instances de classes qui ne doivent être instanciées qu'une seule fois (par exemple, les clients de DB).

Adaptateur : utile lorsque nous avons des changements dans une ressource externe, mais que nous ne voulons pas refactoriser tout notre code pour nous adapter à ces changements. Nous créons alors un adaptateur pour créer une compatibility layer.

5. Qu-est-ce que le NewSQL ?

Lorsqu'on choisit une base de données pour un projet, la integrite et l'évolutivité sont généralement deux aspects contradictoires. Les bases de données SQL respectent les principes ACID, mais ne sont pas très performantes en matière d'intégrité des données, par contre que les bases de données NoSQL sont bien meilleures en termes d'évolutivité, au détriment de la cohérence. NewSQL est un type de base de données relationnelle qui vise à offrir une évolutivité tout en conservant les avantages des bases de données SQL.