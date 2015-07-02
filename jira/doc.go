// Package jira implement slackcmd.Commander to handle request about jira ticket,
// it will get ticket IDs from request, and give back ticket details into the channel that sent command.
//
// This package (slackcmd also) use Viper (https://github.com/spf13/viper) for manage configuration of app,
// so, copy config from file `config.yml.dist` to your application config, replace by your config value,
// and use Viper to load the config.
//
// Check out example package in (github.com/ngdinhtoan/slackcmd/example) to see how to use this package.
package jira
