Шаблоны проектирования

1) Фабричный метод (Factory Method)

Назначение: Определяет интерфейс для создания объекта, но оставляет подклассам решение о том, какой класс инстанцировать.

UML:

![diag1.PNG](diag%2Fdiag1.PNG)

Код

// Scenario интерфейс

`type Scenario interface {
CreateScenario()
}`

// TextScenario структура, реализующая интерфейс Scenario

`type TextScenario struct{}`

`func (t TextScenario) CreateScenario() {
fmt.Println("Creating a text scenario")
}`

// AudioScenario структура, реализующая интерфейс Scenario
type AudioScenario struct{}

`func (a AudioScenario) CreateScenario() {
fmt.Println("Creating an audio scenario")
}`

// ScenarioFactory интерфейс

`type ScenarioFactory interface {
CreateScenario() Scenario
}`

// TextScenarioFactory структура, реализующая интерфейс ScenarioFactory

`type TextScenarioFactory struct{}`

`func (t TextScenarioFactory) CreateScenario() Scenario {
return TextScenario{}
}`

// AudioScenarioFactory структура, реализующая интерфейс ScenarioFactory

`type AudioScenarioFactory struct{}`

`func (a AudioScenarioFactory) CreateScenario() Scenario
{
    return AudioScenario{}
}`

// Пример использования
`func main() {`

    var textFactory ScenarioFactory = TextScenarioFactory{}
    textScenario := textFactory.CreateScenario()
    textScenario.CreateScenario()

	var audioFactory ScenarioFactory = AudioScenarioFactory{}
	audioScenario := audioFactory.CreateScenario()
	audioScenario.CreateScenario()
`}`

Объяснение

Интерфейс Scenario определяет метод CreateScenario(). Его реализуют структуры TextScenario и AudioScenario,
каждая из которых предоставляет свою реализацию метода CreateScenario(), выводя соответствующее сообщение о создании сценария.

Интерфейс фабрики ScenarioFactory определяет метод CreateScenario(), возвращающий объект типа Scenario. 
Его реализуют структуры TextScenarioFactory и AudioScenarioFactory, каждая из которых создает соответствующий сценарий (TextScenario или AudioScenario).

TextScenario и AudioScenario реализуют интерфейс Scenario, поэтому метод фабрики ScenarioFactory не обязан реализовывать каждый тип сценария, а только интерфейс Scenario. Конкретные реализации уже определяются в подклассах фабрик.

2) Абстрактная фабрика (Abstract Factory)

Назначение: Предоставляет интерфейс для создания семейств взаимосвязанных или взаимозависимых объектов без указания их конкретных классов.

UML:

![diag2.PNG](diag%2Fdiag2.PNG)

Код

// Интерфейсы учебных материалов

`type GrammarLesson interface {
Study()
}`

`type VocabularyExercise interface {
Practice()
}`

// Интерфейс фабрики

`type MarineEnglishFactory interface {
CreateGrammarLesson() GrammarLesson
CreateVocabularyExercise() VocabularyExercise
}`

// Реализации для начинающих

`type BeginnerGrammarLesson struct{}`

`func (b BeginnerGrammarLesson) Study() {
fmt.Println("Изучаем базовую морскую грамматику: части корабля и команды")
}`

`type BeginnerVocabularyExercise struct{}`

`func (b BeginnerVocabularyExercise) Practice() {
fmt.Println("Практикуем базовую лексику: левый борт, правый борт, якорь")
}`

// Реализации для продвинутых

`type AdvancedGrammarLesson struct{}`

`func (a AdvancedGrammarLesson) Study() {
fmt.Println("Осваиваем сложную грамматику: протоколы радиосвязи")
}`

`type AdvancedVocabularyExercise struct{}`

`func (a AdvancedVocabularyExercise) Practice() {
fmt.Println("Практикуем продвинутую терминологию: метеосводки и навигационные термины")
}`

// Конкретные фабрики

`type BeginnerMarineFactory struct{}`

`func (b BeginnerMarineFactory) CreateGrammarLesson() GrammarLesson {
return BeginnerGrammarLesson{}
}`

`func (b BeginnerMarineFactory) CreateVocabularyExercise() VocabularyExercise {
return BeginnerVocabularyExercise{}
}`

`type AdvancedMarineFactory struct{}`

`func (a AdvancedMarineFactory) CreateGrammarLesson() GrammarLesson {
return AdvancedGrammarLesson{}
}`

`func (a AdvancedMarineFactory) CreateVocabularyExercise() VocabularyExercise {
return AdvancedVocabularyExercise{}
}`

// Клиент обучения

`type EnglishLearner struct {
lesson   GrammarLesson
exercise VocabularyExercise
}`

`func NewEnglishLearner(factory MarineEnglishFactory) *EnglishLearner {
return &EnglishLearner{
lesson:   factory.CreateGrammarLesson(),
exercise: factory.CreateVocabularyExercise(),
}
}`

`func (e *EnglishLearner) PrepareForStudy() {
fmt.Println("\nПодготовка к занятию по морскому английскому:")
e.lesson.Study()
e.exercise.Practice()
}`

// Пример использования

`func main() {
beginnerFactory := BeginnerMarineFactory{}
advancedFactory := AdvancedMarineFactory{}`

	beginnerLearner := NewEnglishLearner(beginnerFactory)
	beginnerLearner.PrepareForStudy()

	advancedLearner := NewEnglishLearner(advancedFactory)
	advancedLearner.PrepareForStudy()
`}
`
Описание

Вводятся два интерфейса: GrammarLesson и VocabularyExercise, которые определяют методы Study() и Practice() соответственно. Создаются конкретные структуры, реализующие эти интерфейсы. В данном случае, BeginnerGrammarLesson и AdvancedGrammarLesson — это два варианта уроков грамматики, а BeginnerVocabularyExercise и AdvancedVocabularyExercise — два варианта упражнений на лексику. Каждая из этих структур реализует соответствующие методы интерфейсов.

Создается интерфейс MarineEnglishFactory, предоставляющий интерфейс для создания связанных объектов уроков грамматики и упражнений на лексику. Создаются две конкретные реализации абстрактной фабрики: BeginnerMarineFactory и AdvancedMarineFactory. Каждая из этих фабрик реализует методы CreateGrammarLesson и CreateVocabularyExercise, возвращая соответствующие объекты уроков и упражнений.

Есть класс клиент EnglishLearner, который не содержит конкретные объекты уроков и упражнений (может быть любой) и объекта фабрики (может быть любая). В зависимости от использования класса клиента будет определяться, какие именно у него уроки и упражнения.

3) Одиночка (Singleton)

Назначение: Гарантирует, что класс имеет только один экземпляр и предоставляет к нему глобальную точку доступа.

UML

![diag3.PNG](diag%2Fdiag3.PNG)

Код

// MarineEnglishStudy структура

`type MarineEnglishStudy struct{}`

// экземпляр MarineEnglishStudy

`var instance *MarineEnglishStudy`

// функция для получения единственного экземпляра MarineEnglishStudy

`func GetMarineEnglishStudy() *MarineEnglishStudy {
if instance == nil {
instance = &MarineEnglishStudy{}
}
return instance
}`

// метод для начала обучения

`func (m *MarineEnglishStudy) StartLearning() {
fmt.Println("Starting to learn Marine English")
}`

// Пример использования

`func main() {
studyInstance := GetMarineEnglishStudy()
studyInstance.StartLearning()
}`

Описание

Есть структура MarineEnglishStudy, в ней есть статическая переменная instance, которая хранит единственный экземпляр данной структуры. Инициализатор функции GetMarineEnglishStudy приватный, чтобы не могли создаваться другие экземпляры.

Метод StartLearning используется для начала процесса обучения.

Структурные шаблоны
1. Адаптер (Adapter)

Назначение: Позволяет объектам с несовместимыми интерфейсами работать вместе.

UML

![diag4.PNG](diag%2Fdiag4.PNG)

Код

// ScenarioGenerator - интерфейс для генерации учебных сценариев

`type ScenarioGenerator interface {
GenerateScenario()
}`

// MarineMaterial - существующий материал для адаптации

`type MarineMaterial struct{}`

`func (m MarineMaterial) LoadContent() {
fmt.Println("Loading nautical vocabulary and protocols")
}`

// MaterialToScenarioAdapter - адаптирует материал в учебный сценарий

`type MaterialToScenarioAdapter struct {
material MarineMaterial
}`

`func NewAdapter(m MarineMaterial) *MaterialToScenarioAdapter {
return &MaterialToScenarioAdapter{material: m}
}`

`func (a *MaterialToScenarioAdapter) GenerateScenario() {
fmt.Println("Adapting marine content to learning scenario...")
a.material.LoadContent()
fmt.Println("Scenario created: Bridge communication exercise")
}`

// Пример использования

`func main() {
nauticalContent := MarineMaterial{}
learningAdapter := NewAdapter(nauticalContent)`

	// Генерация сценария через адаптер
	learningAdapter.GenerateScenario()
`}`

Описание

Интерфейс ScenarioGenerator определяет метод GenerateScenario для генерации учебных сценариев.

Создается структура MarineMaterial, представляющая существующий материал для адаптации. Метод LoadContent этой структуры загружает морскую лексику и протоколы, выводя соответствующее сообщение.

Определяется структура MaterialToScenarioAdapter, которая адаптирует материал в учебный сценарий. В ней хранится экземпляр MarineMaterial. Функция NewAdapter создает новый адаптер, принимая MarineMaterial в качестве аргумента.

Метод GenerateScenario структуры MaterialToScenarioAdapter адаптирует морской контент в учебный сценарий, загружая содержимое через метод LoadContent и выводя сообщение о создании сценария, например, "Упражнение по связи на мостике".

2. Декоратор (Decorator)

Назначение: Позволяет динамически добавлять новые обязанности объектам.

UML

![diag5.PNG](diag%2Fdiag5.PNG)

Код


// LearningModule - интерфейс для учебных модулей

`type LearningModule interface {
ExecuteScenario()
}`

// BasicNavigationModule - базовый модуль обучения

`type BasicNavigationModule struct{}`

`func (b *BasicNavigationModule) ExecuteScenario() {
fmt.Println("Executing basic navigation communication scenario")
}`

// ScenarioDecorator - декоратор для расширения сценариев

`type ScenarioDecorator struct {
module LearningModule
}`

`func NewScenarioDecorator(module LearningModule) *ScenarioDecorator {
return &ScenarioDecorator{module: module}
}`

`func (s *ScenarioDecorator) ExecuteScenario() {
s.module.ExecuteScenario()
s.addAdvancedScenario()
}`

`func (s *ScenarioDecorator) addAdvancedScenario() {
fmt.Println("Adding emergency procedures scenario")
}`

// Пример использования

`func main() {
basicModule := &BasicNavigationModule{}
advancedCourse := NewScenarioDecorator(basicModule)`

	fmt.Println("Basic course:")
	basicModule.ExecuteScenario()

	fmt.Println("\nAdvanced course:")
	advancedCourse.ExecuteScenario()
`}`

Описание

Интерфейс LearningModule определяет метод ExecuteScenario для выполнения учебных сценариев.

Создается структура BasicNavigationModule, представляющая базовый модуль обучения. Метод ExecuteScenario этой структуры выполняет сценарий базовой навигационной связи, выводя соответствующее сообщение.

Определяется структура ScenarioDecorator, которая служит декоратором для расширения сценариев. В ней хранится экземпляр LearningModule. Функция NewScenarioDecorator создает новый декоратор, принимая модуль обучения в качестве аргумента.

Метод ExecuteScenario структуры ScenarioDecorator сначала вызывает метод ExecuteScenario у обернутого модуля, а затем добавляет дополнительный сценарий, выводя сообщение о добавлении сценария по экстренным процедурам.

3. Компоновщик (Composite)

Назначение: Позволяет объединять объекты в древовидные структуры для представления иерархий "часть-целое".

UML:

![diag6.PNG](diag%2Fdiag6.PNG)

Код

// TrainingMaterial - интерфейс для учебных материалов

`type TrainingMaterial interface {
ExecuteScenario()
}`

// Exercise - конкретное упражнение

`type Exercise struct {
Title string
}`

`func NewExercise(title string) *Exercise {
return &Exercise{Title: title}
}`

`func (e *Exercise) ExecuteScenario() {
fmt.Printf("Executing exercise: %s\n", e.Title)
}`

// ScenarioComposite - составной сценарий

`type ScenarioComposite struct {
materials []TrainingMaterial
}`

`func NewScenarioComposite() *ScenarioComposite {
return &ScenarioComposite{
materials: make([]TrainingMaterial, 0),
}
}`

`func (s *ScenarioComposite) AddMaterial(material TrainingMaterial) {
s.materials = append(s.materials, material)
}`

`func (s *ScenarioComposite) RemoveMaterial(material TrainingMaterial) {
for i, m := range s.materials {
if m == material {
s.materials = append(s.materials[:i], s.materials[i+1:]...)
break
}
}
}`

`func (s *ScenarioComposite) ExecuteScenario() {
fmt.Println("Starting composite scenario:")
for _, material := range s.materials {
material.ExecuteScenario()
}
fmt.Println("Scenario completed\n")
}`

`func main() {`

// Создание упражнений

`radioCommunication := NewExercise("Radio phraseology practice")
emergencyDrill := NewExercise("Emergency situation drill")
navigationTest := NewExercise("Navigation commands test")`

	// Базовый сценарий
	basicScenario := NewScenarioComposite()
	basicScenario.AddMaterial(radioCommunication)
	basicScenario.AddMaterial(navigationTest)

	// Продвинутый сценарий
	advancedScenario := NewScenarioComposite()
	advancedScenario.AddMaterial(basicScenario)
	advancedScenario.AddMaterial(emergencyDrill)

	// Выполнение сценариев
	fmt.Println("Basic training:")
	basicScenario.ExecuteScenario()

	fmt.Println("Advanced training:")
	advancedScenario.ExecuteScenario()
`}`

Описание

Интерфейс TrainingMaterial определяет метод ExecuteScenario для выполнения учебных материалов.

Создается структура Exercise, представляющая конкретное упражнение. Функция NewExercise создает новый экземпляр Exercise с заданным заголовком. Метод ExecuteScenario этой структуры выполняет упражнение, выводя сообщение с его заголовком.

Определяется структура ScenarioComposite, которая представляет составной сценарий. В ней хранится срез материалов типа TrainingMaterial. Функция NewScenarioComposite создает новый экземпляр ScenarioComposite с пустым срезом материалов.

Методы AddMaterial и RemoveMaterial позволяют добавлять и удалять учебные материалы из составного сценария. Метод ExecuteScenario запускает выполнение всех сценариев, хранящихся в materials, выводя сообщение о начале и завершении составного сценария.

4. Фасад (Facade)

Назначение: Предоставляет упрощенный интерфейс к сложной системе классов, библиотек или фреймворков.

UML:

![diag7.PNG](diag%2Fdiag7.PNG)

Код

// MaterialCreator - класс для создания учебных материалов

`type MaterialCreator struct{}`

`func (m *MaterialCreator) CreateMaterial() {
fmt.Println("Creating training material for marine English.")
}`

// ScenarioCreator - класс для создания сценариев

`type ScenarioCreator struct{}`

`func (s *ScenarioCreator) CreateScenario() {
fmt.Println("Creating training scenario for marine communication.")
}`

// TrainingFacade - фасад для упрощения создания материалов и сценариев

`type TrainingFacade struct {
materialCreator *MaterialCreator
scenarioCreator *ScenarioCreator
}`

`func NewTrainingFacade() *TrainingFacade {
return &TrainingFacade{
materialCreator: &MaterialCreator{},
scenarioCreator: &ScenarioCreator{},
}
}`

`func (f *TrainingFacade) CreateTraining() {
f.materialCreator.CreateMaterial()
f.scenarioCreator.CreateScenario()
}`

// Пример использования

`func main() {
facade := NewTrainingFacade()
facade.CreateTraining()
}`

Описание

Класс MaterialCreator отвечает за создание учебных материалов. Метод CreateMaterial выводит сообщение о создании учебного материала для морского английского.

Создается класс ScenarioCreator, который отвечает за создание сценариев. Метод CreateScenario выводит сообщение о создании учебного сценария для морской коммуникации.

Определяется класс TrainingFacade, который служит фасадом для упрощения процесса создания материалов и сценариев. В нем хранятся экземпляры MaterialCreator и ScenarioCreator. Функция NewTrainingFacade создает новый экземпляр TrainingFacade, инициализируя его создателями материалов и сценариев.

Метод CreateTraining класса TrainingFacade вызывает методы создания материалов и сценариев, обеспечивая единый интерфейс для выполнения этих операций.

Поведенческие шаблоны
1. Цепочка обязанностей (Chain of Responsibility)

Назначение: Позволяет передавать запросы по цепочке обработчиков, где каждый обработчик решает, обрабатывать запрос или передать его дальше.

UML:

![diag8.PNG](diag%2Fdiag8.PNG)

Код

// Handler - абстрактный класс для обработки запросов

`type Handler interface {
SetSuccessor(successor Handler)
HandleRequest(request string)
}`

// BaseHandler - базовая структура для обработки запросов

`type BaseHandler struct {
successor Handler
}`

`func (h *BaseHandler) SetSuccessor(successor Handler) {
h.successor = successor
}`

// MaterialHandler - обработчик для учебных материалов

`type MaterialHandler struct {
BaseHandler
}
`
`func (h *MaterialHandler) HandleRequest(request string) {
if request == "Material" {
fmt.Println("MaterialHandler: Handling material request.")
} else if h.successor != nil {
h.successor.HandleRequest(request)
}
}`

// ScenarioHandler - обработчик для сценариев

`type ScenarioHandler struct {
BaseHandler
}`

`func (h *ScenarioHandler) HandleRequest(request string) {
if request == "Scenario" {
fmt.Println("ScenarioHandler: Handling scenario request.")
} else if h.successor != nil {
h.successor.HandleRequest(request)
}
}`

// Пример использования

`func main() {
materialHandler := &MaterialHandler{}
scenarioHandler := &ScenarioHandler{}`

	// Устанавливаем цепочку обработчиков
	materialHandler.SetSuccessor(scenarioHandler)

	// Обработка запросов
	materialHandler.HandleRequest("Material") // Обработает запрос на материал
	materialHandler.HandleRequest("Scenario")  // Обработает запрос на сценарий
	materialHandler.HandleRequest("Other")     // Не будет обработан
`}`

Описание

Интерфейс Handler определяет методы для обработки запросов. Метод SetSuccessor устанавливает следующего обработчика в цепочке, а метод HandleRequest обрабатывает входящий запрос.

Создается базовая структура BaseHandler, которая реализует интерфейс Handler. В ней хранится ссылка на следующего обработчика. Метод SetSuccessor устанавливает значение для этого поля.

Определяется структура MaterialHandler, которая наследует BaseHandler и реализует метод HandleRequest. Если запрос равен "Material", обработчик выполняет соответствующее действие. Если запрос не соответствует, он передает его следующему обработчику в цепочке, если таковой установлен.

Создается структура ScenarioHandler, которая также наследует BaseHandler и реализует метод HandleRequest. Аналогично, если запрос равен "Scenario", обработчик выполняет действие, а в противном случае передает запрос следующему обработчику.

2. Команда (Command)

Назначение: Инкапсулирует запрос как объект, позволяя параметризовать клиентов с различными запросами, ставить запросы в очередь и поддерживать отмену операций.

UML:

![diag9.PNG](diag%2Fdiag9.PNG)

Код

// Command - интерфейс для команд

`type Command interface {
Execute()
}`

// MaterialReceiver - класс, который выполняет действия

`type MaterialReceiver struct{}`

`func (r *MaterialReceiver) Action() {
fmt.Println("MaterialReceiver: Creating training material for marine English.")
}`

// CreateMaterialCommand - конкретная команда для создания материалов

`type CreateMaterialCommand struct {
receiver *MaterialReceiver
}`

`func NewCreateMaterialCommand(receiver *MaterialReceiver) *CreateMaterialCommand {
return &CreateMaterialCommand{receiver: receiver}
}`

`func (c *CreateMaterialCommand) Execute() {
c.receiver.Action()
}`

// Invoker - класс, который вызывает команды

`type Invoker struct {
command Command
}`

`func (i *Invoker) SetCommand(command Command) {
i.command = command
}`

`func (i *Invoker) Invoke() {
i.command.Execute()
}`

// Пример использования

`func main() {
receiver := &MaterialReceiver{}
command := NewCreateMaterialCommand(receiver)
invoker := &Invoker{}
invoker.SetCommand(command)
invoker.Invoke()
}`

Описание

Интерфейс Command определяет метод Execute для выполнения команд.

Создается класс MaterialReceiver, который выполняет действия. Метод Action этого класса выводит сообщение о создании учебного материала для морского английского.

Определяется структура CreateMaterialCommand, представляющая конкретную команду для создания материалов. В ней хранится ссылка на MaterialReceiver. Функция NewCreateMaterialCommand создает новый экземпляр CreateMaterialCommand, принимая MaterialReceiver в качестве аргумента. Метод Execute вызывает метод Action у получателя.

Создается класс Invoker, который отвечает за вызов команд. В нем хранится команда типа Command. Метод SetCommand устанавливает команду, а метод Invoke выполняет команду, вызывая ее метод Execute.

3. Итератор (Iterator)

Назначение: Предоставляет способ последовательного доступа к элементам агрегированного объекта, не раскрывая его внутреннего представления.

UML:

![diag10.PNG](diag%2Fdiag10.PNG)

Код

// Iterator - интерфейс для итераторов

`type Iterator interface {
HasNext() bool
Next() interface{}
}`

// Aggregate - интерфейс для агрегатов

`type Aggregate interface {
CreateIterator() Iterator
}`

// MaterialAggregate - агрегат для учебных материалов

`type MaterialAggregate struct {
materials []string
}`

`func (a *MaterialAggregate) CreateIterator() Iterator {
return &MaterialIterator{aggregate: a, current: 0}
}
`
// MaterialIterator - итератор для учебных материалов

`type MaterialIterator struct {
aggregate *MaterialAggregate
current   int
}`

`func (i *MaterialIterator) HasNext() bool {
return i.current < len(i.aggregate.materials)
}`

`func (i *MaterialIterator) Next() interface{} {
item := i.aggregate.materials[i.current]
i.current++
return item
}`

// Пример использования

`func main() {
aggregate := &MaterialAggregate{
materials: []string{"Material 1", "Material 2", "Material 3"},
}`

	iterator := aggregate.CreateIterator()
	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}
`}`

Описание

Интерфейс Iterator определяет методы HasNext для проверки наличия следующего элемента и Next для получения следующего элемента.

Создается интерфейс Aggregate, который определяет метод CreateIterator для создания итератора.

Определяется структура MaterialAggregate, представляющая агрегат для учебных материалов. В ней хранится срез строк, представляющий материалы. Метод CreateIterator создает и возвращает новый экземпляр MaterialIterator, инициализируя его текущим индексом.

Создается структура MaterialIterator, которая реализует интерфейс Iterator. В ней хранится ссылка на MaterialAggregate и текущий индекс. Метод HasNext проверяет, есть ли еще элементы в агрегате, а метод Next возвращает текущий элемент и увеличивает индекс.

4. Наблюдатель (Observer)

Назначение: Определяет зависимость "один ко многим" между объектами, так что при изменении состояния одного объекта все его зависимые объекты уведомляются и обновляются автоматически.

UML:

![diag11.PNG](diag%2Fdiag11.PNG)

Код


// Observer - интерфейс для наблюдателей

`type Observer interface {
Update()
}`

// Subject - интерфейс для субъекта

`type Subject interface {
Attach(observer Observer)
Detach(observer Observer)
Notify()
}`

// MaterialSubject - конкретный субъект, который уведомляет наблюдателей

`type MaterialSubject struct {
observers []Observer
}`

`func (s *MaterialSubject) Attach(observer Observer) {
s.observers = append(s.observers, observer)
}`

`func (s *MaterialSubject) Detach(observer Observer) {
for i, obs := range s.observers {
if obs == observer {
s.observers = append(s.observers[:i], s.observers[i+1:]...)
break
}
}
}`

`func (s *MaterialSubject) Notify() {
for _, observer := range s.observers {
observer.Update()
}
}`

// MaterialObserver - конкретный наблюдатель, который реагирует на изменения

`type MaterialObserver struct {
subject *MaterialSubject
}`

`func NewMaterialObserver(subject *MaterialSubject) *MaterialObserver {
return &MaterialObserver{subject: subject}
}`

`func (o *MaterialObserver) Update() {
fmt.Println("MaterialObserver: Received update from subject.")
}`

// Пример использования

`func main() {
subject := &MaterialSubject{}
observer := NewMaterialObserver(subject)`

	subject.Attach(observer)
	subject.Notify()
`}`

Описание

Интерфейс Observer определяет метод Update для обновления наблюдателей.

Создается интерфейс Subject, который определяет методы для управления наблюдателями: Attach для добавления наблюдателя, Detach для удаления наблюдателя и Notify для уведомления всех наблюдателей об изменениях.

Определяется структура MaterialSubject, представляющая конкретный субъект, который уведомляет наблюдателей. В ней хранится срез наблюдателей. Метод Attach добавляет нового наблюдателя в список, а метод Detach удаляет наблюдателя из списка. Метод Notify вызывает метод Update у всех зарегистрированных наблюдателей.

Создается структура MaterialObserver, представляющая конкретного наблюдателя, который реагирует на изменения в субъекте. В ней хранится ссылка на MaterialSubject. Функция NewMaterialObserver создает новый экземпляр MaterialObserver, принимая субъект в качестве аргумента. Метод Update выводит сообщение о получении обновления от субъекта.

5. Состояние (State)

Назначение: Позволяет объекту изменять свое поведение при изменении его внутреннего состояния. Объект будет казаться изменившим свой класс.

UML:

![diag12.PNG](diag%2Fdiag12.PNG)

Код:

// State - интерфейс для состояний

`type State interface {
Handle(context *Context)
}`

// MaterialStateA - конкретное состояние A

`type MaterialStateA struct{}`

`func (s *MaterialStateA) Handle(context *Context) {
fmt.Println("Handling in Material State A")
context.SetState(&MaterialStateB{})
}`

// MaterialStateB - конкретное состояние B

`type MaterialStateB struct{}`

`func (s *MaterialStateB) Handle(context *Context) {
fmt.Println("Handling in Material State B")
context.SetState(&MaterialStateA{})
}`

// Context - контекст, который управляет состоянием

`type Context struct {
state State
}`

`func (c *Context) SetState(state State) {
c.state = state
}`

`func (c *Context) Request() {
c.state.Handle(c)
}`

// Пример использования

`func main() {
context := &Context{state: &MaterialStateA{}}
context.Request() // Handling in Material State A
context.Request() // Handling in Material State B
}`

Описание

Интерфейс State определяет метод Handle, принимающий контекст в качестве аргумента и обрабатывающий состояние.

Создается структура MaterialStateA, представляющая конкретное состояние A. Метод Handle этой структуры выводит сообщение о том, что обрабатывается состояние A, и изменяет состояние контекста на MaterialStateB.

Определяется структура MaterialStateB, представляющая конкретное состояние B. Метод Handle этой структуры выводит сообщение о том, что обрабатывается состояние B, и изменяет состояние контекста на MaterialStateA.

Создается структура Context, которая управляет текущим состоянием. В ней хранится текущее состояние типа State. Метод SetState устанавливает новое состояние, а метод Request вызывает метод Handle у текущего состояния, передавая контекст.