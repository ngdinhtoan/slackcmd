// Package jira implement slackcmd.Commander to handle request about jira ticket,
// it will get ticket IDs from request, and give back ticket details into the channel that sent command.
//
// This package (and package slackcmd also) use Viper (https://github.com/spf13/viper) to manage configuration of app.
// So, copy config from file `config.yml.dist` to your application config, replace by your config value,
// and use Viper to load the config.
//
// Check out package example (github.com/ngdinhtoan/slackcmd/example) to see how to use this package.
package jira
