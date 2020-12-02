package search

var wordMax = 65

/*
func FindRefs(r campwiz.Result, props map[string]*campwiz.Property) []*campwiz.Campground {
	var matching []*campwiz.Campground

	for _, prop := range props {
		for _, cg := range prop.Campgrounds {
			if r.ResID
		for _, sid := range xref. {
			if sid == r.ID {
				matching = append(matching, xref)
				continue
			}
		}
	}

	variations := []string{
		r.Name,
		strings.Join(strings.Split(mangle.Shortest(mangle.Expand(r.Name)), " "), ""),
		mangle.Shortest(r.Name),
		mangle.Expand(r.Name),
		mangle.Shortest(mangle.Expand(r.Name)),
	}

	klog.V(2).Infof("Merge Variations: %v", strings.Join(variations, "|"))

	for _, name := range variations {
		mm := fuzzyMatch(name, xrefs)
		if len(mm) == 0 {
			continue
		}

		if len(mm) == 1 {
			return mm
		}

		if len(mm) > 1 {
			// So, we have multiple matches. Perhaps the locale will help? We no longer have it :(
			// BETTER IDEA: Fuzzy coordinates match?
			klog.V(2).Infof("No unique for %s: %+v - returning all", name, mm)
			return mm
		}
	}

	return matching
}

// fuzzyMatch finds the most likely matching cross-references for a site by name
func fuzzyMatch(name string, xrefs map[string]campwiz.Ref) []campwiz.Campground {
	if name == "" {
		klog.Warningf("fuzzyMatch called with empty name")
		return nil
	}

	keyName := strings.ToUpper(name)
	klog.V(1).Infof("fuzzyMatch(%s) ...", keyName)

	// Three levels of matches.
	var exact []campwiz.Campground
	var prefix []campwiz.Campground
	var contains []campwiz.Campground
	var allWords []campwiz.Campground
	var someWords []campwiz.Campground
	var singleWord []campwiz.Campground

	keywords := strings.Split(keyName, " ")

	for _, xref := range xrefs {
		k := xref.Name
		i := strings.Index(strings.ToLower(k), strings.ToLower(keyName))
		klog.V(4).Infof("Testing: keyName=%s == k=%s (index=%d)", keyName, k, i)
		// The whole key does not exist.
		if i == -1 {
			var wordMatches []string
			kwords := strings.Split(k, " ")
			for _, kw := range kwords {
				for _, keyword := range keywords {
					if keyword == kw {
						wordMatches = append(wordMatches, kw)
					}
				}
			}
			if len(wordMatches) == len(keywords) {
				klog.V(2).Infof("All words match for %s: %s", keyName, k)
				allWords = append(allWords, xref)
			} else if len(wordMatches) > 1 {
				klog.V(2).Infof("Partial match for %s: %s (matches=%v)", keyName, k, wordMatches)
				someWords = append(someWords, xref)
			} else if len(wordMatches) == 1 {
				klog.V(3).Infof("Found single word match for %s: %s (matches=%v)", keyName, k, wordMatches)
				singleWord = append(singleWord, xref)
			}
			continue
		}
		if i == 0 {
			if strings.HasPrefix(k, keyName+" - ") {
				exact = append(exact, xref)
				klog.V(2).Infof("Found exact match for %s: %s", keyName, k)
				continue
			}
			klog.V(2).Infof("Found prefix match for %s: %s", keyName, k)
			prefix = append(prefix, xref)
			continue
		} else if i > 0 {
			klog.V(2).Infof("Found substring match for %s: %s", keyName, k)
			contains = append(contains, xref)
		}
	}

	if len(exact) > 0 {
		return exact
	}
	if len(prefix) > 0 {
		return prefix
	}
	if len(contains) > 0 {
		return contains
	}
	if len(allWords) > 0 {
		return allWords
	}
	if len(someWords) > 0 {
		return someWords
	}
	if len(singleWord) == 1 {
		return singleWord
	}
	return nil
}

*/
