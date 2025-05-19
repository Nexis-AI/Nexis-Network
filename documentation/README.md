# Nexis Network Documentation

This directory contains high level notes about the Nexis Network codebase.
It briefly describes some components and suggests future enhancements
for decentralized AI compute.

## Overview

Nexis Network is derived from the Solana blockchain project. It contains
multiple workspace crates providing runtime, validator and client
functionality.

## AI Compute Roadmap

To enable node operators to dedicate processing power for AI workloads,
new modules can be introduced. The new `ai-node` crate (see `ai-node/`)
illustrates a small TCP service that will evolve into a compute worker.
Future development can expand upon this by:

- Integrating job scheduling and workload distribution
- Supporting model downloads from Hugging Face
- Providing APIs for Vercel AI SDK or Supabase integration

These additions will pave the way for decentralized model serving and
training across the Nexis Network.
