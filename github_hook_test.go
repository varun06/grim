package grim

// Copyright 2015 MediaMath <http://www.mediamath.com>.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"fmt"
	"testing"
)

func TestPushHook(t *testing.T) {
	hook, err := extractHookEvent(pushBody)
	if err != nil {
		t.Fatal(err)
	}

	expected := hookEvent{
		eventName: "push",
		action:    "",
		userName:  "bhand-mm",
		owner:     "MediaMath",
		repo:      "grim",
		target:    "test",
		ref:       "ade10d0a64f122d095e1b33cdb5719099f542288",
		statusRef: "ade10d0a64f122d095e1b33cdb5719099f542288",
		url:       "https://github.com/MediaMath/grim/compare/d6bc37a5a405...ade10d0a64f1",
		prNumber:  0,
	}

	failIfDifferent(t, *hook, expected)
}

func TestPullRequestHook(t *testing.T) {
	hook, err := extractHookEvent(prBody)
	if err != nil {
		t.Fatal(err)
	}

	expected := hookEvent{
		eventName: "pull_request",
		action:    "reopened",
		userName:  "bhand-mm",
		owner:     "MediaMath",
		repo:      "grim",
		target:    "master",
		ref:       "",
		statusRef: "566f52c6f30600abe63cd43ffbb74a2da30dba68",
		url:       "https://github.com/MediaMath/grim/pull/34",
		prNumber:  34,
	}

	failIfDifferent(t, *hook, expected)
}

func failIfDifferent(t *testing.T, first, second hookEvent) {
	firstStr := fmt.Sprintf("%+v", first)
	secondStr := fmt.Sprintf("%+v", second)
	if firstStr != secondStr {
		t.Fatalf("hooks don't match: %v != %v", firstStr, secondStr)
	}
}

var (
	pushBody = `{
  "Type" : "Notification",
  "MessageId" : "37c7ca8e-9401-5d53-95d9-873fadf38796",
  "TopicArn" : "arn:aws:sns:us-east-1:888665229551:platform-infra-changelog-github-hooks",
  "Message" : "{\"ref\":\"refs/heads/test\",\"before\":\"d6bc37a5a405055db0edbe616671b50f2a9db1df\",\"after\":\"ade10d0a64f122d095e1b33cdb5719099f542288\",\"created\":false,\"deleted\":false,\"forced\":false,\"base_ref\":null,\"compare\":\"https://github.com/MediaMath/grim/compare/d6bc37a5a405...ade10d0a64f1\",\"commits\":[{\"id\":\"ade10d0a64f122d095e1b33cdb5719099f542288\",\"distinct\":true,\"message\":\"test\",\"timestamp\":\"2015-04-22T00:54:54-05:00\",\"url\":\"https://github.com/MediaMath/grim/commit/ade10d0a64f122d095e1b33cdb5719099f542288\",\"author\":{\"name\":\"Billy Hand\",\"email\":\"bhand@mediamath.com\",\"username\":\"bhand-mm\"},\"committer\":{\"name\":\"Billy Hand\",\"email\":\"bhand@mediamath.com\",\"username\":\"bhand-mm\"},\"added\":[],\"removed\":[],\"modified\":[\"test\"]}],\"head_commit\":{\"id\":\"ade10d0a64f122d095e1b33cdb5719099f542288\",\"distinct\":true,\"message\":\"test\",\"timestamp\":\"2015-04-22T00:54:54-05:00\",\"url\":\"https://github.com/MediaMath/grim/commit/ade10d0a64f122d095e1b33cdb5719099f542288\",\"author\":{\"name\":\"Billy Hand\",\"email\":\"bhand@mediamath.com\",\"username\":\"bhand-mm\"},\"committer\":{\"name\":\"Billy Hand\",\"email\":\"bhand@mediamath.com\",\"username\":\"bhand-mm\"},\"added\":[],\"removed\":[],\"modified\":[\"test\"]},\"repository\":{\"id\":33742142,\"name\":\"grim\",\"full_name\":\"MediaMath/grim\",\"owner\":{\"name\":\"MediaMath\",\"email\":\"open-source@mediamath.com\"},\"private\":true,\"html_url\":\"https://github.com/MediaMath/grim\",\"description\":\"github responder in mediamath\",\"fork\":false,\"url\":\"https://github.com/MediaMath/grim\",\"forks_url\":\"https://api.github.com/repos/MediaMath/grim/forks\",\"keys_url\":\"https://api.github.com/repos/MediaMath/grim/keys{/key_id}\",\"collaborators_url\":\"https://api.github.com/repos/MediaMath/grim/collaborators{/collaborator}\",\"teams_url\":\"https://api.github.com/repos/MediaMath/grim/teams\",\"hooks_url\":\"https://api.github.com/repos/MediaMath/grim/hooks\",\"issue_events_url\":\"https://api.github.com/repos/MediaMath/grim/issues/events{/number}\",\"events_url\":\"https://api.github.com/repos/MediaMath/grim/events\",\"assignees_url\":\"https://api.github.com/repos/MediaMath/grim/assignees{/user}\",\"branches_url\":\"https://api.github.com/repos/MediaMath/grim/branches{/branch}\",\"tags_url\":\"https://api.github.com/repos/MediaMath/grim/tags\",\"blobs_url\":\"https://api.github.com/repos/MediaMath/grim/git/blobs{/sha}\",\"git_tags_url\":\"https://api.github.com/repos/MediaMath/grim/git/tags{/sha}\",\"git_refs_url\":\"https://api.github.com/repos/MediaMath/grim/git/refs{/sha}\",\"trees_url\":\"https://api.github.com/repos/MediaMath/grim/git/trees{/sha}\",\"statuses_url\":\"https://api.github.com/repos/MediaMath/grim/statuses/{sha}\",\"languages_url\":\"https://api.github.com/repos/MediaMath/grim/languages\",\"stargazers_url\":\"https://api.github.com/repos/MediaMath/grim/stargazers\",\"contributors_url\":\"https://api.github.com/repos/MediaMath/grim/contributors\",\"subscribers_url\":\"https://api.github.com/repos/MediaMath/grim/subscribers\",\"subscription_url\":\"https://api.github.com/repos/MediaMath/grim/subscription\",\"commits_url\":\"https://api.github.com/repos/MediaMath/grim/commits{/sha}\",\"git_commits_url\":\"https://api.github.com/repos/MediaMath/grim/git/commits{/sha}\",\"comments_url\":\"https://api.github.com/repos/MediaMath/grim/comments{/number}\",\"issue_comment_url\":\"https://api.github.com/repos/MediaMath/grim/issues/comments{/number}\",\"contents_url\":\"https://api.github.com/repos/MediaMath/grim/contents/{+path}\",\"compare_url\":\"https://api.github.com/repos/MediaMath/grim/compare/{base}...{head}\",\"merges_url\":\"https://api.github.com/repos/MediaMath/grim/merges\",\"archive_url\":\"https://api.github.com/repos/MediaMath/grim/{archive_format}{/ref}\",\"downloads_url\":\"https://api.github.com/repos/MediaMath/grim/downloads\",\"issues_url\":\"https://api.github.com/repos/MediaMath/grim/issues{/number}\",\"pulls_url\":\"https://api.github.com/repos/MediaMath/grim/pulls{/number}\",\"milestones_url\":\"https://api.github.com/repos/MediaMath/grim/milestones{/number}\",\"notifications_url\":\"https://api.github.com/repos/MediaMath/grim/notifications{?since,all,participating}\",\"labels_url\":\"https://api.github.com/repos/MediaMath/grim/labels{/name}\",\"releases_url\":\"https://api.github.com/repos/MediaMath/grim/releases{/id}\",\"created_at\":1428688026,\"updated_at\":\"2015-04-20T20:48:25Z\",\"pushed_at\":1429682098,\"git_url\":\"git://github.com/MediaMath/grim.git\",\"ssh_url\":\"git@github.com:MediaMath/grim.git\",\"clone_url\":\"https://github.com/MediaMath/grim.git\",\"svn_url\":\"https://github.com/MediaMath/grim\",\"homepage\":null,\"size\":329,\"stargazers_count\":0,\"watchers_count\":0,\"language\":\"Go\",\"has_issues\":true,\"has_downloads\":true,\"has_wiki\":true,\"has_pages\":false,\"forks_count\":1,\"mirror_url\":null,\"open_issues_count\":2,\"forks\":1,\"open_issues\":2,\"watchers\":0,\"default_branch\":\"master\",\"stargazers\":0,\"master_branch\":\"master\",\"organization\":\"MediaMath\"},\"pusher\":{\"name\":\"bhand-mm\",\"email\":\"bhand@mediamath.com\"},\"organization\":{\"login\":\"MediaMath\",\"id\":2982134,\"url\":\"https://api.github.com/orgs/MediaMath\",\"repos_url\":\"https://api.github.com/orgs/MediaMath/repos\",\"events_url\":\"https://api.github.com/orgs/MediaMath/events\",\"members_url\":\"https://api.github.com/orgs/MediaMath/members{/member}\",\"public_members_url\":\"https://api.github.com/orgs/MediaMath/public_members{/member}\",\"avatar_url\":\"https://avatars.githubusercontent.com/u/2982134?v=3\",\"description\":\"Performance Reimagined. Marketing Reengineered.\"},\"sender\":{\"login\":\"bhand-mm\",\"id\":5913552,\"avatar_url\":\"https://avatars.githubusercontent.com/u/5913552?v=3\",\"gravatar_id\":\"\",\"url\":\"https://api.github.com/users/bhand-mm\",\"html_url\":\"https://github.com/bhand-mm\",\"followers_url\":\"https://api.github.com/users/bhand-mm/followers\",\"following_url\":\"https://api.github.com/users/bhand-mm/following{/other_user}\",\"gists_url\":\"https://api.github.com/users/bhand-mm/gists{/gist_id}\",\"starred_url\":\"https://api.github.com/users/bhand-mm/starred{/owner}{/repo}\",\"subscriptions_url\":\"https://api.github.com/users/bhand-mm/subscriptions\",\"organizations_url\":\"https://api.github.com/users/bhand-mm/orgs\",\"repos_url\":\"https://api.github.com/users/bhand-mm/repos\",\"events_url\":\"https://api.github.com/users/bhand-mm/events{/privacy}\",\"received_events_url\":\"https://api.github.com/users/bhand-mm/received_events\",\"type\":\"User\",\"site_admin\":false}}",
  "Timestamp" : "2015-04-22T05:54:58.631Z",
  "SignatureVersion" : "1",
  "Signature" : "mYhkydMaMKmntofvlHgGSZCzYqhGv3QEYJ30waakXpzkptjekIseZoVTwFwDuhiiNpeu/Nit2iZUl6qWCTknDMkaAn92U5PsertbK+Uy6JKHErmieSWaUlwU8VpZhov5qND9t+I+29xFMwAwin4t95QOd2dOcz6FRkG7ZhGsq3QC/oq67hU2oSD5nj/Rv1jDySBq7afvTOp3mgdlBr4wQ1zrWwOHOCwRYrAdGpHb+gjbgkUXqQujmPgjlD/Rma6RpLWTiWPJYr9xumOyLVsmcKUdVMgYEAD3of5xNm9vXPBps/0riHOT8jFLaMHdyHQJ6arj32QsuuZxQS4TjMHNKA==",
  "SigningCertURL" : "https://sns.us-east-1.amazonaws.com/SimpleNotificationService-d6d679a1d18e95c2f9ffcf11f4f9e198.pem",
  "UnsubscribeURL" : "https://sns.us-east-1.amazonaws.com/?Action=Unsubscribe&SubscriptionArn=arn:aws:sns:us-east-1:888665229551:platform-infra-changelog-github-hooks:8b7c84bf-d3a8-4a9b-9cdc-48c83cdf5b23"
}`
	prBody = `{
  "Type" : "Notification",
  "MessageId" : "aab19963-dc66-54a2-a1d3-d4d8b6585531",
  "TopicArn" : "arn:aws:sns:us-east-1:888665229551:grim-MediaMath-grim-repo-topic",
  "Message" : "{\"action\":\"reopened\",\"number\":34,\"pull_request\":{\"url\":\"https://api.github.com/repos/MediaMath/grim/pulls/34\",\"id\":34411226,\"html_url\":\"https://github.com/MediaMath/grim/pull/34\",\"diff_url\":\"https://github.com/MediaMath/grim/pull/34.diff\",\"patch_url\":\"https://github.com/MediaMath/grim/pull/34.patch\",\"issue_url\":\"https://api.github.com/repos/MediaMath/grim/issues/34\",\"number\":34,\"state\":\"open\",\"locked\":false,\"title\":\"merge commit sha polling\",\"user\":{\"login\":\"bhand-mm\",\"id\":5913552,\"avatar_url\":\"https://avatars.githubusercontent.com/u/5913552?v=3\",\"gravatar_id\":\"\",\"url\":\"https://api.github.com/users/bhand-mm\",\"html_url\":\"https://github.com/bhand-mm\",\"followers_url\":\"https://api.github.com/users/bhand-mm/followers\",\"following_url\":\"https://api.github.com/users/bhand-mm/following{/other_user}\",\"gists_url\":\"https://api.github.com/users/bhand-mm/gists{/gist_id}\",\"starred_url\":\"https://api.github.com/users/bhand-mm/starred{/owner}{/repo}\",\"subscriptions_url\":\"https://api.github.com/users/bhand-mm/subscriptions\",\"organizations_url\":\"https://api.github.com/users/bhand-mm/orgs\",\"repos_url\":\"https://api.github.com/users/bhand-mm/repos\",\"events_url\":\"https://api.github.com/users/bhand-mm/events{/privacy}\",\"received_events_url\":\"https://api.github.com/users/bhand-mm/received_events\",\"type\":\"User\",\"site_admin\":false},\"body\":\"\",\"created_at\":\"2015-04-29T22:07:47Z\",\"updated_at\":\"2015-04-29T23:23:42Z\",\"closed_at\":null,\"merged_at\":null,\"merge_commit_sha\":\"485342cd439a0db2d33a8b4bafe01467d1e25e4a\",\"assignee\":null,\"milestone\":null,\"commits_url\":\"https://api.github.com/repos/MediaMath/grim/pulls/34/commits\",\"review_comments_url\":\"https://api.github.com/repos/MediaMath/grim/pulls/34/comments\",\"review_comment_url\":\"https://api.github.com/repos/MediaMath/grim/pulls/comments{/number}\",\"comments_url\":\"https://api.github.com/repos/MediaMath/grim/issues/34/comments\",\"statuses_url\":\"https://api.github.com/repos/MediaMath/grim/statuses/566f52c6f30600abe63cd43ffbb74a2da30dba68\",\"head\":{\"label\":\"bhand-mm:master\",\"ref\":\"master\",\"sha\":\"566f52c6f30600abe63cd43ffbb74a2da30dba68\",\"user\":{\"login\":\"bhand-mm\",\"id\":5913552,\"avatar_url\":\"https://avatars.githubusercontent.com/u/5913552?v=3\",\"gravatar_id\":\"\",\"url\":\"https://api.github.com/users/bhand-mm\",\"html_url\":\"https://github.com/bhand-mm\",\"followers_url\":\"https://api.github.com/users/bhand-mm/followers\",\"following_url\":\"https://api.github.com/users/bhand-mm/following{/other_user}\",\"gists_url\":\"https://api.github.com/users/bhand-mm/gists{/gist_id}\",\"starred_url\":\"https://api.github.com/users/bhand-mm/starred{/owner}{/repo}\",\"subscriptions_url\":\"https://api.github.com/users/bhand-mm/subscriptions\",\"organizations_url\":\"https://api.github.com/users/bhand-mm/orgs\",\"repos_url\":\"https://api.github.com/users/bhand-mm/repos\",\"events_url\":\"https://api.github.com/users/bhand-mm/events{/privacy}\",\"received_events_url\":\"https://api.github.com/users/bhand-mm/received_events\",\"type\":\"User\",\"site_admin\":false},\"repo\":{\"id\":33742432,\"name\":\"grim\",\"full_name\":\"bhand-mm/grim\",\"owner\":{\"login\":\"bhand-mm\",\"id\":5913552,\"avatar_url\":\"https://avatars.githubusercontent.com/u/5913552?v=3\",\"gravatar_id\":\"\",\"url\":\"https://api.github.com/users/bhand-mm\",\"html_url\":\"https://github.com/bhand-mm\",\"followers_url\":\"https://api.github.com/users/bhand-mm/followers\",\"following_url\":\"https://api.github.com/users/bhand-mm/following{/other_user}\",\"gists_url\":\"https://api.github.com/users/bhand-mm/gists{/gist_id}\",\"starred_url\":\"https://api.github.com/users/bhand-mm/starred{/owner}{/repo}\",\"subscriptions_url\":\"https://api.github.com/users/bhand-mm/subscriptions\",\"organizations_url\":\"https://api.github.com/users/bhand-mm/orgs\",\"repos_url\":\"https://api.github.com/users/bhand-mm/repos\",\"events_url\":\"https://api.github.com/users/bhand-mm/events{/privacy}\",\"received_events_url\":\"https://api.github.com/users/bhand-mm/received_events\",\"type\":\"User\",\"site_admin\":false},\"private\":true,\"html_url\":\"https://github.com/bhand-mm/grim\",\"description\":\"github responder in mediamath\",\"fork\":true,\"url\":\"https://api.github.com/repos/bhand-mm/grim\",\"forks_url\":\"https://api.github.com/repos/bhand-mm/grim/forks\",\"keys_url\":\"https://api.github.com/repos/bhand-mm/grim/keys{/key_id}\",\"collaborators_url\":\"https://api.github.com/repos/bhand-mm/grim/collaborators{/collaborator}\",\"teams_url\":\"https://api.github.com/repos/bhand-mm/grim/teams\",\"hooks_url\":\"https://api.github.com/repos/bhand-mm/grim/hooks\",\"issue_events_url\":\"https://api.github.com/repos/bhand-mm/grim/issues/events{/number}\",\"events_url\":\"https://api.github.com/repos/bhand-mm/grim/events\",\"assignees_url\":\"https://api.github.com/repos/bhand-mm/grim/assignees{/user}\",\"branches_url\":\"https://api.github.com/repos/bhand-mm/grim/branches{/branch}\",\"tags_url\":\"https://api.github.com/repos/bhand-mm/grim/tags\",\"blobs_url\":\"https://api.github.com/repos/bhand-mm/grim/git/blobs{/sha}\",\"git_tags_url\":\"https://api.github.com/repos/bhand-mm/grim/git/tags{/sha}\",\"git_refs_url\":\"https://api.github.com/repos/bhand-mm/grim/git/refs{/sha}\",\"trees_url\":\"https://api.github.com/repos/bhand-mm/grim/git/trees{/sha}\",\"statuses_url\":\"https://api.github.com/repos/bhand-mm/grim/statuses/{sha}\",\"languages_url\":\"https://api.github.com/repos/bhand-mm/grim/languages\",\"stargazers_url\":\"https://api.github.com/repos/bhand-mm/grim/stargazers\",\"contributors_url\":\"https://api.github.com/repos/bhand-mm/grim/contributors\",\"subscribers_url\":\"https://api.github.com/repos/bhand-mm/grim/subscribers\",\"subscription_url\":\"https://api.github.com/repos/bhand-mm/grim/subscription\",\"commits_url\":\"https://api.github.com/repos/bhand-mm/grim/commits{/sha}\",\"git_commits_url\":\"https://api.github.com/repos/bhand-mm/grim/git/commits{/sha}\",\"comments_url\":\"https://api.github.com/repos/bhand-mm/grim/comments{/number}\",\"issue_comment_url\":\"https://api.github.com/repos/bhand-mm/grim/issues/comments{/number}\",\"contents_url\":\"https://api.github.com/repos/bhand-mm/grim/contents/{+path}\",\"compare_url\":\"https://api.github.com/repos/bhand-mm/grim/compare/{base}...{head}\",\"merges_url\":\"https://api.github.com/repos/bhand-mm/grim/merges\",\"archive_url\":\"https://api.github.com/repos/bhand-mm/grim/{archive_format}{/ref}\",\"downloads_url\":\"https://api.github.com/repos/bhand-mm/grim/downloads\",\"issues_url\":\"https://api.github.com/repos/bhand-mm/grim/issues{/number}\",\"pulls_url\":\"https://api.github.com/repos/bhand-mm/grim/pulls{/number}\",\"milestones_url\":\"https://api.github.com/repos/bhand-mm/grim/milestones{/number}\",\"notifications_url\":\"https://api.github.com/repos/bhand-mm/grim/notifications{?since,all,participating}\",\"labels_url\":\"https://api.github.com/repos/bhand-mm/grim/labels{/name}\",\"releases_url\":\"https://api.github.com/repos/bhand-mm/grim/releases{/id}\",\"created_at\":\"2015-04-10T17:52:59Z\",\"updated_at\":\"2015-04-29T23:07:01Z\",\"pushed_at\":\"2015-04-29T23:07:01Z\",\"git_url\":\"git://github.com/bhand-mm/grim.git\",\"ssh_url\":\"git@github.com:bhand-mm/grim.git\",\"clone_url\":\"https://github.com/bhand-mm/grim.git\",\"svn_url\":\"https://github.com/bhand-mm/grim\",\"homepage\":null,\"size\":353,\"stargazers_count\":0,\"watchers_count\":0,\"language\":\"Go\",\"has_issues\":false,\"has_downloads\":true,\"has_wiki\":true,\"has_pages\":false,\"forks_count\":0,\"mirror_url\":null,\"open_issues_count\":0,\"forks\":0,\"open_issues\":0,\"watchers\":0,\"default_branch\":\"master\"}},\"base\":{\"label\":\"MediaMath:master\",\"ref\":\"master\",\"sha\":\"a9aa7476fb09fc09a9a2bbd246f6191165ffd772\",\"user\":{\"login\":\"MediaMath\",\"id\":2982134,\"avatar_url\":\"https://avatars.githubusercontent.com/u/2982134?v=3\",\"gravatar_id\":\"\",\"url\":\"https://api.github.com/users/MediaMath\",\"html_url\":\"https://github.com/MediaMath\",\"followers_url\":\"https://api.github.com/users/MediaMath/followers\",\"following_url\":\"https://api.github.com/users/MediaMath/following{/other_user}\",\"gists_url\":\"https://api.github.com/users/MediaMath/gists{/gist_id}\",\"starred_url\":\"https://api.github.com/users/MediaMath/starred{/owner}{/repo}\",\"subscriptions_url\":\"https://api.github.com/users/MediaMath/subscriptions\",\"organizations_url\":\"https://api.github.com/users/MediaMath/orgs\",\"repos_url\":\"https://api.github.com/users/MediaMath/repos\",\"events_url\":\"https://api.github.com/users/MediaMath/events{/privacy}\",\"received_events_url\":\"https://api.github.com/users/MediaMath/received_events\",\"type\":\"Organization\",\"site_admin\":false},\"repo\":{\"id\":33742142,\"name\":\"grim\",\"full_name\":\"MediaMath/grim\",\"owner\":{\"login\":\"MediaMath\",\"id\":2982134,\"avatar_url\":\"https://avatars.githubusercontent.com/u/2982134?v=3\",\"gravatar_id\":\"\",\"url\":\"https://api.github.com/users/MediaMath\",\"html_url\":\"https://github.com/MediaMath\",\"followers_url\":\"https://api.github.com/users/MediaMath/followers\",\"following_url\":\"https://api.github.com/users/MediaMath/following{/other_user}\",\"gists_url\":\"https://api.github.com/users/MediaMath/gists{/gist_id}\",\"starred_url\":\"https://api.github.com/users/MediaMath/starred{/owner}{/repo}\",\"subscriptions_url\":\"https://api.github.com/users/MediaMath/subscriptions\",\"organizations_url\":\"https://api.github.com/users/MediaMath/orgs\",\"repos_url\":\"https://api.github.com/users/MediaMath/repos\",\"events_url\":\"https://api.github.com/users/MediaMath/events{/privacy}\",\"received_events_url\":\"https://api.github.com/users/MediaMath/received_events\",\"type\":\"Organization\",\"site_admin\":false},\"private\":true,\"html_url\":\"https://github.com/MediaMath/grim\",\"description\":\"github responder in mediamath\",\"fork\":false,\"url\":\"https://api.github.com/repos/MediaMath/grim\",\"forks_url\":\"https://api.github.com/repos/MediaMath/grim/forks\",\"keys_url\":\"https://api.github.com/repos/MediaMath/grim/keys{/key_id}\",\"collaborators_url\":\"https://api.github.com/repos/MediaMath/grim/collaborators{/collaborator}\",\"teams_url\":\"https://api.github.com/repos/MediaMath/grim/teams\",\"hooks_url\":\"https://api.github.com/repos/MediaMath/grim/hooks\",\"issue_events_url\":\"https://api.github.com/repos/MediaMath/grim/issues/events{/number}\",\"events_url\":\"https://api.github.com/repos/MediaMath/grim/events\",\"assignees_url\":\"https://api.github.com/repos/MediaMath/grim/assignees{/user}\",\"branches_url\":\"https://api.github.com/repos/MediaMath/grim/branches{/branch}\",\"tags_url\":\"https://api.github.com/repos/MediaMath/grim/tags\",\"blobs_url\":\"https://api.github.com/repos/MediaMath/grim/git/blobs{/sha}\",\"git_tags_url\":\"https://api.github.com/repos/MediaMath/grim/git/tags{/sha}\",\"git_refs_url\":\"https://api.github.com/repos/MediaMath/grim/git/refs{/sha}\",\"trees_url\":\"https://api.github.com/repos/MediaMath/grim/git/trees{/sha}\",\"statuses_url\":\"https://api.github.com/repos/MediaMath/grim/statuses/{sha}\",\"languages_url\":\"https://api.github.com/repos/MediaMath/grim/languages\",\"stargazers_url\":\"https://api.github.com/repos/MediaMath/grim/stargazers\",\"contributors_url\":\"https://api.github.com/repos/MediaMath/grim/contributors\",\"subscribers_url\":\"https://api.github.com/repos/MediaMath/grim/subscribers\",\"subscription_url\":\"https://api.github.com/repos/MediaMath/grim/subscription\",\"commits_url\":\"https://api.github.com/repos/MediaMath/grim/commits{/sha}\",\"git_commits_url\":\"https://api.github.com/repos/MediaMath/grim/git/commits{/sha}\",\"comments_url\":\"https://api.github.com/repos/MediaMath/grim/comments{/number}\",\"issue_comment_url\":\"https://api.github.com/repos/MediaMath/grim/issues/comments{/number}\",\"contents_url\":\"https://api.github.com/repos/MediaMath/grim/contents/{+path}\",\"compare_url\":\"https://api.github.com/repos/MediaMath/grim/compare/{base}...{head}\",\"merges_url\":\"https://api.github.com/repos/MediaMath/grim/merges\",\"archive_url\":\"https://api.github.com/repos/MediaMath/grim/{archive_format}{/ref}\",\"downloads_url\":\"https://api.github.com/repos/MediaMath/grim/downloads\",\"issues_url\":\"https://api.github.com/repos/MediaMath/grim/issues{/number}\",\"pulls_url\":\"https://api.github.com/repos/MediaMath/grim/pulls{/number}\",\"milestones_url\":\"https://api.github.com/repos/MediaMath/grim/milestones{/number}\",\"notifications_url\":\"https://api.github.com/repos/MediaMath/grim/notifications{?since,all,participating}\",\"labels_url\":\"https://api.github.com/repos/MediaMath/grim/labels{/name}\",\"releases_url\":\"https://api.github.com/repos/MediaMath/grim/releases{/id}\",\"created_at\":\"2015-04-10T17:47:06Z\",\"updated_at\":\"2015-04-29T20:20:30Z\",\"pushed_at\":\"2015-04-29T23:07:01Z\",\"git_url\":\"git://github.com/MediaMath/grim.git\",\"ssh_url\":\"git@github.com:MediaMath/grim.git\",\"clone_url\":\"https://github.com/MediaMath/grim.git\",\"svn_url\":\"https://github.com/MediaMath/grim\",\"homepage\":null,\"size\":429,\"stargazers_count\":2,\"watchers_count\":2,\"language\":\"Go\",\"has_issues\":true,\"has_downloads\":true,\"has_wiki\":true,\"has_pages\":false,\"forks_count\":2,\"mirror_url\":null,\"open_issues_count\":3,\"forks\":2,\"open_issues\":3,\"watchers\":2,\"default_branch\":\"master\"}},\"_links\":{\"self\":{\"href\":\"https://api.github.com/repos/MediaMath/grim/pulls/34\"},\"html\":{\"href\":\"https://github.com/MediaMath/grim/pull/34\"},\"issue\":{\"href\":\"https://api.github.com/repos/MediaMath/grim/issues/34\"},\"comments\":{\"href\":\"https://api.github.com/repos/MediaMath/grim/issues/34/comments\"},\"review_comments\":{\"href\":\"https://api.github.com/repos/MediaMath/grim/pulls/34/comments\"},\"review_comment\":{\"href\":\"https://api.github.com/repos/MediaMath/grim/pulls/comments{/number}\"},\"commits\":{\"href\":\"https://api.github.com/repos/MediaMath/grim/pulls/34/commits\"},\"statuses\":{\"href\":\"https://api.github.com/repos/MediaMath/grim/statuses/566f52c6f30600abe63cd43ffbb74a2da30dba68\"}},\"merged\":false,\"mergeable\":null,\"mergeable_state\":\"unknown\",\"merged_by\":null,\"comments\":0,\"review_comments\":0,\"commits\":8,\"additions\":118,\"deletions\":28,\"changed_files\":5},\"repository\":{\"id\":33742142,\"name\":\"grim\",\"full_name\":\"MediaMath/grim\",\"owner\":{\"login\":\"MediaMath\",\"id\":2982134,\"avatar_url\":\"https://avatars.githubusercontent.com/u/2982134?v=3\",\"gravatar_id\":\"\",\"url\":\"https://api.github.com/users/MediaMath\",\"html_url\":\"https://github.com/MediaMath\",\"followers_url\":\"https://api.github.com/users/MediaMath/followers\",\"following_url\":\"https://api.github.com/users/MediaMath/following{/other_user}\",\"gists_url\":\"https://api.github.com/users/MediaMath/gists{/gist_id}\",\"starred_url\":\"https://api.github.com/users/MediaMath/starred{/owner}{/repo}\",\"subscriptions_url\":\"https://api.github.com/users/MediaMath/subscriptions\",\"organizations_url\":\"https://api.github.com/users/MediaMath/orgs\",\"repos_url\":\"https://api.github.com/users/MediaMath/repos\",\"events_url\":\"https://api.github.com/users/MediaMath/events{/privacy}\",\"received_events_url\":\"https://api.github.com/users/MediaMath/received_events\",\"type\":\"Organization\",\"site_admin\":false},\"private\":true,\"html_url\":\"https://github.com/MediaMath/grim\",\"description\":\"github responder in mediamath\",\"fork\":false,\"url\":\"https://api.github.com/repos/MediaMath/grim\",\"forks_url\":\"https://api.github.com/repos/MediaMath/grim/forks\",\"keys_url\":\"https://api.github.com/repos/MediaMath/grim/keys{/key_id}\",\"collaborators_url\":\"https://api.github.com/repos/MediaMath/grim/collaborators{/collaborator}\",\"teams_url\":\"https://api.github.com/repos/MediaMath/grim/teams\",\"hooks_url\":\"https://api.github.com/repos/MediaMath/grim/hooks\",\"issue_events_url\":\"https://api.github.com/repos/MediaMath/grim/issues/events{/number}\",\"events_url\":\"https://api.github.com/repos/MediaMath/grim/events\",\"assignees_url\":\"https://api.github.com/repos/MediaMath/grim/assignees{/user}\",\"branches_url\":\"https://api.github.com/repos/MediaMath/grim/branches{/branch}\",\"tags_url\":\"https://api.github.com/repos/MediaMath/grim/tags\",\"blobs_url\":\"https://api.github.com/repos/MediaMath/grim/git/blobs{/sha}\",\"git_tags_url\":\"https://api.github.com/repos/MediaMath/grim/git/tags{/sha}\",\"git_refs_url\":\"https://api.github.com/repos/MediaMath/grim/git/refs{/sha}\",\"trees_url\":\"https://api.github.com/repos/MediaMath/grim/git/trees{/sha}\",\"statuses_url\":\"https://api.github.com/repos/MediaMath/grim/statuses/{sha}\",\"languages_url\":\"https://api.github.com/repos/MediaMath/grim/languages\",\"stargazers_url\":\"https://api.github.com/repos/MediaMath/grim/stargazers\",\"contributors_url\":\"https://api.github.com/repos/MediaMath/grim/contributors\",\"subscribers_url\":\"https://api.github.com/repos/MediaMath/grim/subscribers\",\"subscription_url\":\"https://api.github.com/repos/MediaMath/grim/subscription\",\"commits_url\":\"https://api.github.com/repos/MediaMath/grim/commits{/sha}\",\"git_commits_url\":\"https://api.github.com/repos/MediaMath/grim/git/commits{/sha}\",\"comments_url\":\"https://api.github.com/repos/MediaMath/grim/comments{/number}\",\"issue_comment_url\":\"https://api.github.com/repos/MediaMath/grim/issues/comments{/number}\",\"contents_url\":\"https://api.github.com/repos/MediaMath/grim/contents/{+path}\",\"compare_url\":\"https://api.github.com/repos/MediaMath/grim/compare/{base}...{head}\",\"merges_url\":\"https://api.github.com/repos/MediaMath/grim/merges\",\"archive_url\":\"https://api.github.com/repos/MediaMath/grim/{archive_format}{/ref}\",\"downloads_url\":\"https://api.github.com/repos/MediaMath/grim/downloads\",\"issues_url\":\"https://api.github.com/repos/MediaMath/grim/issues{/number}\",\"pulls_url\":\"https://api.github.com/repos/MediaMath/grim/pulls{/number}\",\"milestones_url\":\"https://api.github.com/repos/MediaMath/grim/milestones{/number}\",\"notifications_url\":\"https://api.github.com/repos/MediaMath/grim/notifications{?since,all,participating}\",\"labels_url\":\"https://api.github.com/repos/MediaMath/grim/labels{/name}\",\"releases_url\":\"https://api.github.com/repos/MediaMath/grim/releases{/id}\",\"created_at\":\"2015-04-10T17:47:06Z\",\"updated_at\":\"2015-04-29T20:20:30Z\",\"pushed_at\":\"2015-04-29T23:07:01Z\",\"git_url\":\"git://github.com/MediaMath/grim.git\",\"ssh_url\":\"git@github.com:MediaMath/grim.git\",\"clone_url\":\"https://github.com/MediaMath/grim.git\",\"svn_url\":\"https://github.com/MediaMath/grim\",\"homepage\":null,\"size\":429,\"stargazers_count\":2,\"watchers_count\":2,\"language\":\"Go\",\"has_issues\":true,\"has_downloads\":true,\"has_wiki\":true,\"has_pages\":false,\"forks_count\":2,\"mirror_url\":null,\"open_issues_count\":3,\"forks\":2,\"open_issues\":3,\"watchers\":2,\"default_branch\":\"master\"},\"organization\":{\"login\":\"MediaMath\",\"id\":2982134,\"url\":\"https://api.github.com/orgs/MediaMath\",\"repos_url\":\"https://api.github.com/orgs/MediaMath/repos\",\"events_url\":\"https://api.github.com/orgs/MediaMath/events\",\"members_url\":\"https://api.github.com/orgs/MediaMath/members{/member}\",\"public_members_url\":\"https://api.github.com/orgs/MediaMath/public_members{/member}\",\"avatar_url\":\"https://avatars.githubusercontent.com/u/2982134?v=3\",\"description\":\"Performance Reimagined. Marketing Reengineered.\"},\"sender\":{\"login\":\"bhand-mm\",\"id\":5913552,\"avatar_url\":\"https://avatars.githubusercontent.com/u/5913552?v=3\",\"gravatar_id\":\"\",\"url\":\"https://api.github.com/users/bhand-mm\",\"html_url\":\"https://github.com/bhand-mm\",\"followers_url\":\"https://api.github.com/users/bhand-mm/followers\",\"following_url\":\"https://api.github.com/users/bhand-mm/following{/other_user}\",\"gists_url\":\"https://api.github.com/users/bhand-mm/gists{/gist_id}\",\"starred_url\":\"https://api.github.com/users/bhand-mm/starred{/owner}{/repo}\",\"subscriptions_url\":\"https://api.github.com/users/bhand-mm/subscriptions\",\"organizations_url\":\"https://api.github.com/users/bhand-mm/orgs\",\"repos_url\":\"https://api.github.com/users/bhand-mm/repos\",\"events_url\":\"https://api.github.com/users/bhand-mm/events{/privacy}\",\"received_events_url\":\"https://api.github.com/users/bhand-mm/received_events\",\"type\":\"User\",\"site_admin\":false}}",
  "Timestamp" : "2015-04-29T23:23:42.593Z",
  "SignatureVersion" : "1",
  "Signature" : "tBd6aHsnkWCC1RocDYVNHrUkBCn8ih2WucPIMhuY2hgbgXcYDW3u9maRnQFCR5yileVE9l8i7kSpvoL9JJhRJuPnktunxeGWrjO2Ir8+P0fSSSTe0MOgPLQqyPLOTSGh0rNYxqtU4oVanHZqXYWdJQ7bIiS1sjNydVRUpVL0BCWod+yVZwpyotTpuLLoz7wIAUpWzVXTOs/u7SoNRhf3spWpRNHKyMcf9Uv1yFs6/d9nrNJuqDcM1vW/5X4a3E+LzMN4WB2wAjvB+xIluXH2bgI/0eqy0JcqeMnrJh/on0nI/2NrHQs2dRJqWakbh9cGTQXMzaOjlpBsDHv2IoG7Hg==",
  "SigningCertURL" : "https://sns.us-east-1.amazonaws.com/SimpleNotificationService-d6d679a1d18e95c2f9ffcf11f4f9e198.pem",
  "UnsubscribeURL" : "https://sns.us-east-1.amazonaws.com/?Action=Unsubscribe&SubscriptionArn=arn:aws:sns:us-east-1:888665229551:grim-MediaMath-grim-repo-topic:7feefee7-4a04-486d-8d41-fd1b08dbb26c"
}`
)
