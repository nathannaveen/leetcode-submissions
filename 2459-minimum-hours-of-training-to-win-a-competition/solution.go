func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
    extra := 0
    
    for _, e := range energy {
        if initialEnergy <= e {
            extra += (e + 1) - initialEnergy
            initialEnergy = e + 1
        }
        initialEnergy -= e
    }
    
    for _, e := range experience {
        if initialExperience <= e {
            extra += (e + 1) - initialExperience
            initialExperience = e + 1
        }
        initialExperience += e
    }
    
    return extra
}