def latest(scores):
    """Returns the latest score"""

    if len(scores) == 0 or scores is None:
        raise Exception("Score list is empty")

    return scores[len(scores) - 1]


def personal_best(scores):
    """Returns the best score"""

    if len(scores) == 0 or scores is None:
        raise exception("score list is empty")

    return max(scores)


def personal_top_three(scores):
    """Returns the best three scores

    If there is less than three scores, returns the scores sorted descending
    """

    if len(scores) == 0 or scores is None:
        raise Exception("Score list is empty")

    scores.sort(reverse=True)

    if len(scores) < 3:
        return scores

    return scores[0:3]
