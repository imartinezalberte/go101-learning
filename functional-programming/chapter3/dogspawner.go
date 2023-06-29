package chapter3

const (
	Bulldog Breed = iota
	Havanese
	Cavalier
	Poodle

	Male Gender = iota
	Female
)

var (
	mHavaneseSpawner = DogSpawner(Havanese, Male)
	fHavaneseSpawner = DogSpawner(Havanese, Female)

	bulldogSpawner  = DogSpawnerFCurrying(Bulldog)
	havaneseSpawner = DogSpawnerFCurrying(Havanese)
	cavalierSpawner = DogSpawnerFCurrying(Cavalier)
	poodleSpawner   = DogSpawnerFCurrying(Poodle)

	maleBulldogSpawner    = bulldogSpawner(Male)
	femaleBulldogSpawner  = bulldogSpawner(Female)
	maleHavaneseSpawner   = havaneseSpawner(Male)
	femaleHavaneseSpawner = havaneseSpawner(Female)
	maleCavalierSpawner   = cavalierSpawner(Male)
	femaleCavalierSpawner = cavalierSpawner(Female)
	malePoodleSpawner     = poodleSpawner(Male)
	femalePoodleSpawner   = poodleSpawner(Female)

	m = map[string]NameToDogFn{
		"male bulldog":    maleBulldogSpawner,
		"female bulldog":  femaleBulldogSpawner,
		"male havanese":   maleHavaneseSpawner,
		"female havanese": femaleHavaneseSpawner,
		"male cavalier":   maleCavalierSpawner,
		"female cavalier": femaleCavalierSpawner,
		"male poodle":     malePoodleSpawner,
		"female poodle":   femalePoodleSpawner,
	}
)

type (
	Name string

	Breed int

	Gender int

	BreedFn func(Breed) GenderFn

	GenderFn func(Gender) NameToDogFn

	NameToDogFn func(Name) Dog

	Dog struct {
		Name   Name
		Breed  Breed
		Gender Gender
	}
)

func createDogsWithoutPartialApplication() {
	bucky := Dog{
		Name:   "Bucky",
		Breed:  Havanese,
		Gender: Male,
	}

	rocky := Dog{
		Name:   "Rocky",
		Breed:  Havanese,
		Gender: Male,
	}

	tipsy := Dog{
		Name:   "Tipsy",
		Breed:  Poodle,
		Gender: Female,
	}

	_, _, _ = bucky, rocky, tipsy
}

func DogSpawner(breed Breed, gender Gender) NameToDogFn {
	return func(name Name) Dog {
		return Dog{
			Name:   name,
			Breed:  breed,
			Gender: gender,
		}
	}
}

func DogSpawnerFCurrying(breed Breed) GenderFn {
	return func(gender Gender) NameToDogFn {
		return func(name Name) Dog {
			return Dog{
				Name:   name,
				Breed:  breed,
				Gender: gender,
			}
		}
	}
}
